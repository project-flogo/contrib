package timer

import (
	"context"
	"fmt"
	"time"

	"github.com/carlescere/scheduler"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

type HandlerSettings struct {
	StartInterval  string `md:"startDelay"`
	RepeatInterval string `md:"repeatInterval"`
}

var triggerMd = trigger.NewMetadata(&HandlerSettings{})

func init() {
	trigger.Register(&TimerTrigger{}, &Factory{})
}

type Factory struct {
}

// Metadata implements trigger.Factory.Metadata
func (*Factory) Metadata() *trigger.Metadata {
	return triggerMd
}

// New implements trigger.Factory.New
func (*Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	return &TimerTrigger{}, nil
}

type TimerTrigger struct {
	timers   []*scheduler.Job
	handlers []trigger.Handler
	logger   log.Logger
}

// Init implements trigger.Init
func (t *TimerTrigger) Initialize(ctx trigger.InitContext) error {

	t.handlers = ctx.GetHandlers()
	t.logger = ctx.Logger()

	return nil
}

// Start implements ext.TimerTrigger.Start
func (t *TimerTrigger) Start() error {

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

// Stop implements ext.TimerTrigger.Stop
func (t *TimerTrigger) Stop() error {

	for _, timer := range t.timers {

		if timer.IsRunning() {
			timer.Quit <- true
		}
	}

	t.timers = nil

	return nil
}

func (t *TimerTrigger) scheduleOnce(handler trigger.Handler, settings *HandlerSettings) error {

	seconds := 0

	if settings.StartInterval != "" {
		d, err := time.ParseDuration(settings.StartInterval)
		if err != nil {
			return fmt.Errorf("unable to parse start delay: %s", err.Error())
		}

		seconds = int(d.Seconds())
		t.logger.Debugf("Scheduling action to run once in %d seconds", seconds)
	}

	timerJob := scheduler.Every(seconds).Seconds()

	fn := func() {
		t.logger.Debug("Executing \"Once\" timer trigger")

		_, err := handler.Handle(context.Background(), nil)
		if err != nil {
			t.logger.Error("Error running handler: ", err.Error())
		}

		if timerJob != nil {
			timerJob.Quit <- true
		}
	}

	if seconds == 0 {
		t.logger.Debug("Start delay not specified, executing action immediately")
		fn()
	} else {
		timerJob, err := timerJob.NotImmediately().Run(fn)
		if err != nil {
			t.logger.Error("Error scheduling execute \"once\" timer: ", err.Error())
		}

		t.timers = append(t.timers, timerJob)
	}

	return nil
}

func (t *TimerTrigger) scheduleRepeating(handler trigger.Handler, settings *HandlerSettings) error {
	t.logger.Info("Scheduling a repeating timer")

	startSeconds := 0

	if settings.StartInterval != "" {
		d, err := time.ParseDuration(settings.StartInterval)
		if err != nil {
			return fmt.Errorf("unable to parse start delay: %s", err.Error())
		}

		startSeconds = int(d.Seconds())
		t.logger.Debugf("Scheduling action to start in %d seconds", startSeconds)
	}

	d, err := time.ParseDuration(settings.RepeatInterval)
	if err != nil {
		return fmt.Errorf("unable to parse repeat interval: %s", err.Error())
	}

	repeatInterval := int(d.Seconds())
	t.logger.Debugf("Scheduling action to repeat every %d seconds", repeatInterval)

	fn := func() {
		t.logger.Debug("Executing \"Repeating\" timer")

		_, err := handler.Handle(context.Background(), nil)
		if err != nil {
			t.logger.Error("Error running handler: ", err.Error())
		}
	}

	if startSeconds == 0 {
		timerJob, err := scheduler.Every(repeatInterval).Seconds().Run(fn)
		if err != nil {
			t.logger.Error("Error scheduling repeating timer: ", err.Error())
		}

		t.timers = append(t.timers, timerJob)
	} else {

		timerJob := scheduler.Every(startSeconds).Seconds()

		fn2 := func() {
			t.logger.Debug("Executing first run of repeating timer")

			_, err := handler.Handle(context.Background(), nil)
			if err != nil {
				t.logger.Error("Error running handler: ", err.Error())
			}

			if timerJob != nil {
				timerJob.Quit <- true
			}

			timerJob, err := scheduler.Every(repeatInterval).Seconds().NotImmediately().Run(fn)
			if err != nil {
				t.logger.Error("Error scheduling repeating timer: ", err.Error())
			}

			t.timers = append(t.timers, timerJob)
		}

		timerJob, err := timerJob.NotImmediately().Run(fn2)
		if err != nil {
			t.logger.Error("Error scheduling delayed start repeating timer: ", err.Error())
		}

		t.timers = append(t.timers, timerJob)
	}

	return nil
}

//type PrintJob struct {
//	Msg string
//}
//
//func (j *PrintJob) Run() error {
//	t.logger.Debug(j.Msg)
//	return nil
//}
