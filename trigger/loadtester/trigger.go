package loadtester

import (
	"fmt"
	"time"

	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

var triggerMd = trigger.NewMetadata(&Settings{}, &Output{})

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
	s := &Settings{StartDelay: 30, Concurrency: 5, Duration: 120}
	err := metadata.MapToStruct(config.Settings, s, true)
	if err != nil {
		return nil, err
	}

	return &Trigger{id: config.Id, settings: s}, nil
}

type Trigger struct {
	id              string
	settings        *Settings
	handler         trigger.Handler
	logger          log.Logger
	statsAggregator chan *RequesterStats
}

func (t *Trigger) Initialize(ctx trigger.InitContext) error {

	t.logger = ctx.Logger()

	if len(ctx.GetHandlers()) == 0 {
		ctx.Logger().Warnf("No Handlers specified for Load Trigger: %s", t.id)
	}

	t.handler = ctx.GetHandlers()[0]

	if t.settings.Handler == "" {
		return nil
	}

	found := false

	for _, handler := range ctx.GetHandlers() {

		if handler.Name() == t.settings.Handler {
			t.handler = handler
			found = true
			break
		}
	}

	if !found {
		ctx.Logger().Warnf("Handler '%s' not found, using first handler", t.settings.Handler)
	}

	return nil
}

// Stop implements util.Managed.Start
func (t *Trigger) Start() error {

	go t.runLoadTest()
	return nil
}

// Stop implements util.Managed.Stop
func (t *Trigger) Stop() error {
	return nil
}

func (t *Trigger) runLoadTest() {

	fmt.Printf("Starting load test in %d seconds\n", t.settings.StartDelay)
	time.Sleep(time.Duration(t.settings.StartDelay) * time.Second)

	data := &Output{Data: t.settings.Data}

	lt := NewLoadTest(t.settings.Duration, t.settings.Concurrency)
	lt.Run(t.handler, data)
}
