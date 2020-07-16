package rest

import (
	"context"
	"strings"

	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

type HandlerSettings struct {
	Lifecycle string `md:"lifecycle,required,allowed(STARTUP,SHUTDOWN)"`
}

var triggerMd = trigger.NewMetadata(&HandlerSettings{})

func init() {
	_ = trigger.Register(&Trigger{}, &Factory{})
}

type Factory struct {
}

// Metadata implements trigger.Factory.Metadata
func (*Factory) Metadata() *trigger.Metadata {
	return triggerMd
}

// New implements trigger.Factory.New
func (*Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	return &Trigger{id: config.Id}, nil
}

// Trigger REST trigger struct
type Trigger struct {
	id               string
	logger           log.Logger

	startupHandlers  []trigger.Handler
	shutdownHandlers []trigger.Handler
}

func (t *Trigger) Initialize(ctx trigger.InitContext) error {

	t.logger = ctx.Logger()

	// Init handlers
	for _, handler := range ctx.GetHandlers() {

		s := &HandlerSettings{}
		err := metadata.MapToStruct(handler.Settings(), s, true)
		if err != nil {
			return err
		}

		if strings.EqualFold(s.Lifecycle, "STARTUP") {
			t.startupHandlers = append(t.shutdownHandlers, handler)
		}

		if strings.EqualFold(s.Lifecycle, "SHUTDOWN") {
			t.startupHandlers = append(t.shutdownHandlers, handler)
		}
	}

	return nil
}

func (t *Trigger) Start() error {
	return nil
}

// Stop implements util.Managed.Stop
func (t *Trigger) Stop() error {
	return nil
}

func (t *Trigger) OnStartup() {

	for _, handler := range t.startupHandlers {
		_, err := handler.Handle(context.Background(), nil)
		if err != nil {
			t.logger.Debugf("Error handling app startup: %s", err.Error())
		}
	}
}

func (t *Trigger) OnShutdown() {
	for _, handler := range t.shutdownHandlers {
		_, err := handler.Handle(context.Background(), nil)
		if err != nil {
			t.logger.Debugf("Error handling app startup: %s", err.Error())
		}
	}
}
