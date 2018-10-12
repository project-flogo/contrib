package timer

import (
	"context"
	"fmt"
	"time"

	"github.com/carlescere/scheduler"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/logger"
	"github.com/project-flogo/core/trigger"
)

var log = logger.GetLogger("trigger-timer")

type HandlerSettings struct {
	StartInterval  string `md:"startDelay"`
	RepeatInterval string `md:"repeatInterval"`
}

var triggerMd = trigger.NewMetadata(&HandlerSettings{})

func init() {
	trigger.Register(&Trigger{}, &Factory{})
}

type Factory struct {
}

// Metadata implements trigger.Factory.Metadata
func (*Factory) Metadata() *trigger.Metadata {
	return triggerMd
}

// New implements trigger.Factory.New
func (*Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	return &Trigger{}, nil
}

type Trigger struct {
	timers   []*scheduler.Job
	handlers []trigger.Handler
}

// Init implements trigger.Init
func (t *Trigger) Initialize(ctx trigger.InitContext) error {

	t.handlers = ctx.GetHandlers()
	return nil
}

// Start implements ext.Trigger.Start
func (t *Trigger) Start() error {

	handlers := t.handlers

	for _, handler := range handlers {

		s := &HandlerSettings{}
		err := metadata.MapToStruct(handler.Settings(), s, true)
		if err != nil {
			return err
		}

		if s.RepeatInterval == "" {
			t.scheduleOnce(handler, s)
		} else {
			t.scheduleRepeating(handler, s)
		}
	}

	return nil
}

// Stop implements ext.Trigger.Stop
func (t *Trigger) Stop() error {

	for _, timer := range t.timers {

		if timer.IsRunning() {
			timer.Quit <- true
		}
	}

	t.timers = nil

	return nil
}

func (t *Trigger) scheduleOnce(handler trigger.Handler, settings *HandlerSettings) error {

	seconds := 0

	if settings.StartInterval != "" {
		d, err := time.ParseDuration(settings.StartInterval)
		if err != nil {
			return fmt.Errorf("unable to parse start delay: %s", err.Error())
		}

		seconds = int(d.Seconds())
		log.Debugf("Scheduling action to run once in %d seconds", seconds)
	}

	timerJob := scheduler.Every(seconds).Seconds()

	fn := func() {
		log.Debug("Executing \"Once\" timer trigger")

		_, err := handler.Handle(context.Background(), nil)
		if err != nil {
			log.Error("Error running handler: ", err.Error())
		}

		if timerJob != nil {
			timerJob.Quit <- true
		}
	}

	if seconds == 0 {
		log.Debug("Start delay not specified, executing action immediately")
		fn()
	} else {
		timerJob, err := timerJob.NotImmediately().Run(fn)
		if err != nil {
			log.Error("Error scheduling execute \"once\" timer: ", err.Error())
		}

		t.timers = append(t.timers, timerJob)
	}

	return nil
}

func (t *Trigger) scheduleRepeating(handler trigger.Handler, settings *HandlerSettings) error {
	log.Info("Scheduling a repeating timer")

	startSeconds := 0

	if settings.StartInterval != "" {
		d, err := time.ParseDuration(settings.StartInterval)
		if err != nil {
			return fmt.Errorf("unable to parse start delay: %s", err.Error())
		}

		startSeconds = int(d.Seconds())
		log.Debugf("Scheduling action to start in %d seconds", startSeconds)
	}

	d, err := time.ParseDuration(settings.RepeatInterval)
	if err != nil {
		return fmt.Errorf("unable to parse repeat interval: %s", err.Error())
	}

	repeatInterval := int(d.Seconds())
	log.Debugf("Scheduling action to repeat every %d seconds", repeatInterval)

	fn := func() {
		log.Debug("Executing \"Repeating\" timer")

		_, err := handler.Handle(context.Background(), nil)
		if err != nil {
			log.Error("Error running handler: ", err.Error())
		}
	}

	if startSeconds == 0 {
		timerJob, err := scheduler.Every(repeatInterval).Seconds().Run(fn)
		if err != nil {
			log.Error("Error scheduling repeating timer: ", err.Error())
		}

		t.timers = append(t.timers, timerJob)
	} else {

		timerJob := scheduler.Every(startSeconds).Seconds()

		fn2 := func() {
			log.Debug("Executing first run of repeating timer")

			_, err := handler.Handle(context.Background(), nil)
			if err != nil {
				log.Error("Error running handler: ", err.Error())
			}

			if timerJob != nil {
				timerJob.Quit <- true
			}

			timerJob, err := scheduler.Every(repeatInterval).Seconds().NotImmediately().Run(fn)
			if err != nil {
				log.Error("Error scheduling repeating timer: ", err.Error())
			}

			t.timers = append(t.timers, timerJob)
		}

		timerJob, err := timerJob.NotImmediately().Run(fn2)
		if err != nil {
			log.Error("Error scheduling delayed start repeating timer: ", err.Error())
		}

		t.timers = append(t.timers, timerJob)
	}

	return nil
}

type PrintJob struct {
	Msg string
}

func (j *PrintJob) Run() error {
	log.Debug(j.Msg)
	return nil
}
