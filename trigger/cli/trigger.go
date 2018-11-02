package cli

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"reflect"

	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

var triggerMd = trigger.NewMetadata(&HandlerSettings{}, &Output{}, &Reply{})

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
	singleton = &Trigger{config: config}
	return singleton, nil
}

var singleton *Trigger

// Trigger CLI trigger struct
type Trigger struct {
	config       *trigger.Config
	handlerInfos []*handlerInfo
	defHandler   trigger.Handler
	logger       log.Logger
}

type handlerInfo struct {
	Invoke  bool
	handler trigger.Handler
}

// Metadata implements trigger.Trigger.Metadata
func (t *Trigger) Metadata() *trigger.Metadata {
	return triggerMd
}

func (t *Trigger) Initialize(ctx trigger.InitContext) error {

	t.logger = ctx.Logger()

	//level, err := logger.GetLevelForName(config.GetLogLevel())
	//
	//if err == nil {
	//	log.SetLogLevel(level)
	//}

	if len(ctx.GetHandlers()) == 0 {
		return fmt.Errorf("no Handlers found for trigger '%s'", t.config.Id)
	}

	hasDefault := false

	// Init handlers
	for _, handler := range ctx.GetHandlers() {

		s := &HandlerSettings{Command: "default"}
		err := metadata.MapToStruct(handler.Settings(), s, true)
		if err != nil {
			return err
		}

		if s.Default {
			t.defHandler = handler
			hasDefault = true
		}

		aInfo := &handlerInfo{Invoke: false, handler: handler}
		t.handlerInfos = append(t.handlerInfos, aInfo)

		xv := reflect.ValueOf(aInfo).Elem()
		addr := xv.FieldByName("Invoke").Addr().Interface()

		switch ptr := addr.(type) {
		case *bool:
			flag.BoolVar(ptr, s.Command, false, "")
		}
	}

	if !hasDefault && len(t.handlerInfos) > 0 {
		t.defHandler = t.handlerInfos[0].handler
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

func Invoke() (string, error) {

	var args []string
	flag.Parse()

	// if we have additional args (after the cmd name and the flow cmd switch)
	// stuff those into args and pass to Invoke(). The action will only receive the
	// additional args that were intending for the action logic.
	if arg := flag.Args(); len(arg) >= 2 {
		args = flag.Args()[2:]
	}

	for _, info := range singleton.handlerInfos {

		if info.Invoke {
			return singleton.Invoke(info.handler, args)
		}
	}

	return singleton.Invoke(singleton.defHandler, args)
}

func (t *Trigger) Invoke(handler trigger.Handler, args []string) (string, error) {

	t.logger.Infof("invoking handler '%s'", handler)

	data := map[string]interface{}{
		"args": args,
	}

	results, err := handler.Handle(context.Background(), data)

	if err != nil {
		t.logger.Debugf("error: %s", err.Error())
		return "", err
	}

	replyData := results["data"]

	if replyData != nil {
		data, err := json.Marshal(replyData)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}

	return "", nil
}
