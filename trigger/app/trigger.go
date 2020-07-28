package rest

import (
	"context"
	"strconv"
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

	startupHandlers  map[string]trigger.Handler
	shutdownHandlers map[string]trigger.Handler
}

func (t *Trigger) Initialize(ctx trigger.InitContext) error {

	t.logger = ctx.Logger()

	// Init handlers
	id := 0
	t.startupHandlers = make(map[string]trigger.Handler)
	t.shutdownHandlers = make(map[string]trigger.Handler)

	for _, handler := range ctx.GetHandlers() {

		s := &HandlerSettings{}
		err := metadata.MapToStruct(handler.Settings(), s, true)
		if err != nil {
			return err
		}

		name := t.id + "-" + handler.Name()
		if name == "" {
			name = t.id + "-handler-" + strconv.Itoa(id)
		}

		if strings.EqualFold(s.Lifecycle, "STARTUP") {
			t.startupHandlers[name] = handler
		} else if strings.EqualFold(s.Lifecycle, "SHUTDOWN") {
			t.startupHandlers[name] = handler
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

func (t *Trigger) OnStartup() error {

	for name, handler := range t.startupHandlers {
		_, err := handler.Handle(context.Background(), nil)
		if err != nil {
			t.logger.Debugf("Error in app startup handler [%s]: %s", name, err.Error())
			return err
		}
	}

	return nil
}

func (t *Trigger) OnShutdown() error {
	var lastErr error
	for name, handler := range t.shutdownHandlers {
		_, err := handler.Handle(context.Background(), nil)
		if err != nil {
			t.logger.Debugf("Error in app shutdown handler [%s]: %s", name, err.Error())
			lastErr = err
		}
	}

	return lastErr
}
