package eftl

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/project-flogo/core/trigger"
	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/data/metadata"
	condition "github.com/mashling/commons/lib/conditions"
	"github.com/mashling/commons/lib/eftl"
	"github.com/mashling/commons/lib/util"
	opentracing "github.com/opentracing/opentracing-go"
)

const (
	settingURL      = "url"
	settingID       = "id"
	settingUser     = "user"
	settingPassword = "password"
	settingCA       = "ca"
	settingDest     = "dest"
)

var triggerMd = trigger.NewMetadata(&Settings{}, &HandlerSettings{}, &Output{})
var logger log.Logger

func init() {
	trigger.Register(&Trigger{}, &Factory{})
}

type Factory struct {
}

// Metadata implements trigger.Factory.Metadata
func (*Factory) Metadata() *trigger.Metadata {
	return triggerMd
}

// Trigger is a simple EFTL trigger
type Trigger struct {
	metadata   *trigger.Metadata
	runner     action.Runner
	config     *trigger.Config
	handlers   map[string]*OptimizedHandler
	connection *eftl.Connection
	stop       chan bool
}

// New implements trigger.Factory.New
func (*Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	s := &Settings{}
	err := metadata.MapToStruct(config.Settings, s, true)
	if err != nil {
		return nil, err
	}
	return &Trigger{}, nil
}

// Init implements trigger.Init
func (t *Trigger) Initialize(ctx trigger.InitContext) error {
	t.runner = *action.Runner
	t.handlers = t.CreateHandlers()
	return nil
}

// CreateHandlers creates handlers mapped to thier topic
func (t *Trigger) CreateHandlers() map[string]*OptimizedHandler {
	handlers := make(map[string]*OptimizedHandler)

	for _, h := range t.config.Handlers {
		tr := h.Settings[settingDest]
		if tr == nil {
			continue
		}
		dest := tr.(string)

		handler := handlers[dest]
		if handler == nil {
			handler = &OptimizedHandler{}
			handlers[dest] = handler
		}

		if condition := h.Settings[util.Flogo_Trigger_Handler_Setting_Condition]; condition != nil {
			dispatch := &Dispatch{
				actionID:   t.config.Id,
				condition:  condition.(string),
				handlerCfg: h,
			}
			handler.dispatches = append(handler.dispatches, dispatch)
		} else {
			handler.defaultActionID = t.config.Id
			handler.defaultHandlerCfg = h
		}
	}

	return handlers
}


// Start implements ext.Trigger.Start
func (t *Trigger) Start() error {

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	ca := t.config.Settings[settingCA]
	if ca != "" {
		certificate, err := ioutil.ReadFile(ca.(string))
		if err != nil {
			logger.Errorf("can't open certificate", err)
			return err
		}
		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(certificate)
		tlsConfig = &tls.Config{
			RootCAs: pool,
		}
	}

	id := t.config.Settings[settingID]
	user := t.config.Settings[settingUser]
	password := t.config.Settings[settingPassword]
	options := &eftl.Options{
		ClientID:  id.(string),
		Username:  user.(string),
		Password:  password.(string),
		TLSConfig: tlsConfig,
	}

	url := t.config.Settings[settingURL]
	errorsChannel := make(chan error, 1)
	connectVal, err := eftl.Connect(url.(string), options, errorsChannel)
	if err != nil {
		logger.Errorf("connection failed: %s", err)
		return err
	}
	t.connection = connectVal
	messages := make(chan eftl.Message, 1000)
	for dest := range t.handlers {
		matcher := fmt.Sprintf("{\"_dest\":\"%s\"}", dest)
		_, err := t.connection.Subscribe(matcher, "", messages)
		if err != nil {
			logger.Errorf("subscription failed: %s", err)
			return err
		}
	}

	t.stop = make(chan bool, 1)
	go func() {
		for {
			select {
			case message := <-messages:
				value := message["_dest"]
				dest, ok := value.(string)
				if !ok {
					logger.Errorf("dest is required for valid message")
					continue
				}
				handler := t.handlers[dest]
				if handler == nil {
					logger.Errorf("no handler for dest ", dest)
					continue
				}
				value = message["content"]
				content, ok := value.([]byte)
				if !ok {
					content = []byte{}
				}
				t.RunAction(handler, dest, content)
			case err := <-errorsChannel:
				logger.Errorf("connection error: %s", err)
			case <-t.stop:
				return
			}
		}
	}()

	return nil
}

// Stop implements ext.Trigger.Stop
func (t *Trigger) Stop() error {
	if t.connection != nil {
		t.connection.Disconnect()
	}
	if t.stop != nil {
		t.stop <- true
	}
	return nil
}

// RunAction starts a new Process Instance
func (t *Trigger) RunAction(handler *OptimizedHandler, dest string, content []byte) {
	logger.Infof("EFTL Trigger: Received request for id '%s'", t.config.Id)

	span := Span{
		Span: opentracing.StartSpan(dest),
	}
	defer span.Finish()

	span.SetTag("dest", dest)
	span.SetTag("content", string(content))

	replyTo, /*data*/_ := t.constructStartRequest(content, span)

	//to do :::::
	//startAttrs, err := t.metadata.Output(data, false)
	//if err != nil {
	//	span.Error("Error setting up attrs: %v", err)
	//}

	actionURI, handlerCfg := handler.GetActionID(string(content), span)
	//action := action.Get(actionURI)
	actionArray := handlerCfg.Actions
	var actions = &action.Action
	for _, act := range actionArray {
		if (act.Config.Id == actionURI){
			actions = act.Act
			break
		}
	}

	context := trigger.NewHandlerContext(context.Background(), handlerCfg)
	replyData, err := t.runner.RunAction(context, actions, map[string]interface{}{})
	if err != nil {
		span.Error("Error starting action: %v", err)
	}
	logger.Debugf("Ran action: [%s]", actionURI)
	span.SetTag("actionURI", actionURI)

	if replyTo == "" {
		return
	}
	reply, err := util.Marshal(replyData)
	if err != nil {
		span.Error("failed to marshal reply data: %v", err)
		return
	}
	span.SetTag("replyTo", replyTo)
	span.SetTag("reply", string(reply))
	err = t.connection.Publish(eftl.Message{
		"_dest":   replyTo,
		"content": reply,
	})
	if err != nil {
		span.Error("failed to send reply data: %v", err)
	}
}

func (t *Trigger) constructStartRequest(message []byte, span Span) (string, map[string]interface{}) {
	span.SetTag("message", string(message))

	var content map[string]interface{}
	err := util.Unmarshal("", message, &content)
	if err != nil {
		span.Error("Error unmarshaling message %s", err.Error())
	}

	replyTo := ""
	pathParams := make(map[string]string)
	queryParams := make(map[string]string)

	mime := ""
	if value, ok := content[util.MetaMIME].(string); ok {
		mime = value
	}
	if mime == util.MIMEApplicationXML {
		getRoot := func() map[string]interface{} {
			body := content[util.XMLKeyBody]
			if body == nil {
				return nil
			}
			for _, e := range body.([]interface{}) {
				element, ok := e.(map[string]interface{})
				if !ok {
					continue
				}
				name, ok := element[util.XMLKeyType].(string)
				if !ok || name != util.XMLTypeElement {
					continue
				}
				return element
			}
			return nil
		}
		root := getRoot()
		fill := func(target string, params map[string]string) {
			rootBody, ok := root[util.XMLKeyBody].([]interface{})
			if !ok {
				return
			}
			for i, e := range rootBody {
				element, ok := e.(map[string]interface{})
				if !ok {
					continue
				}
				name, ok := element[util.XMLKeyName].(string)
				if !ok || name != target {
					continue
				}
				body := element[util.XMLKeyBody]
				if body == nil {
					continue
				}
				for _, e := range body.([]interface{}) {
					element, ok := e.(map[string]interface{})
					if !ok {
						continue
					}
					typ, ok := element[util.XMLKeyType].(string)
					if !ok || typ != util.XMLTypeElement {
						continue
					}
					params[element["key"].(string)] = element["value"].(string)
				}
				root[util.XMLKeyBody] = rootBody[:i+copy(rootBody[i:], rootBody[i+1:])]
				return
			}
		}

		if value, ok := root["replyTo"].(string); ok {
			replyTo = value
			delete(root, "replyTo")
		}
		fill("pathParams", pathParams)
		fill("queryParams", queryParams)
	} else {
		if value, ok := content["replyTo"].(string); ok {
			replyTo = value
			delete(content, "replyTo")
		}

		if params, ok := content["pathParams"].(map[string]interface{}); ok {
			for k, v := range params {
				if param, ok := v.(string); ok {
					pathParams[k] = param
				}
			}
			delete(content, "pathParams")
		}

		if params, ok := content["queryParams"].(map[string]interface{}); ok {
			for k, v := range params {
				if param, ok := v.(string); ok {
					queryParams[k] = param
				}
			}
			delete(content, "queryParams")
		}
	}

	//ctx := opentracing.ContextWithSpan(context.Background(), span)

	data := map[string]interface{}{
		"params":      pathParams,
		"pathParams":  pathParams,
		"queryParams": queryParams,
		"content":     content,
		//"tracing":     ctx,
	}

	return replyTo, data
}


//OptimizedHandler optimized handler
type OptimizedHandler struct {
	defaultActionID   string
	defaultHandlerCfg *trigger.HandlerConfig
	dispatches        []*Dispatch
}

// Error is for reporting errors
func (s *Span) Error(format string, a ...interface{}) {
	str := fmt.Sprintf(format, a...)
	s.SetTag("error", str)
}

// GetActionID gets the action id of the matched handler
func (h *OptimizedHandler) GetActionID(payload string, span Span) (string, *trigger.HandlerConfig) {
	actionID := ""
	var handlerCfg *trigger.HandlerConfig

	for _, dispatch := range h.dispatches {
		expressionStr := dispatch.condition
		//Get condtion and expression type
		conditionOperation, exprType, err := condition.GetConditionOperationAndExpressionType(expressionStr)

		if err != nil || exprType == condition.EXPR_TYPE_NOT_VALID {
			span.Error("not able parse the condition '%v' mentioned for content based handler. skipping the handler.", expressionStr)
			continue
		}

		logger.Debugf("Expression type: %v", exprType)
		logger.Debugf("conditionOperation.LHS %v", conditionOperation.LHS)
		logger.Debugf("conditionOperation.OperatorInfo %v", conditionOperation.OperatorInfo().Names)
		logger.Debugf("conditionOperation.RHS %v", conditionOperation.RHS)

		//Resolve expression's LHS based on expression type and
		//evaluate the expression
		if exprType == condition.EXPR_TYPE_CONTENT {
			exprResult, err := condition.EvaluateCondition(*conditionOperation, payload)
			if err != nil {
				span.Error("not able evaluate expression - %v with error - %v. skipping the handler.", expressionStr, err)
			}
			if exprResult {
				actionID = dispatch.actionID
				handlerCfg = dispatch.handlerCfg
			}
		} else if exprType == condition.EXPR_TYPE_HEADER {
			span.Error("header expression type is invalid for eftl trigger condition")
		} else if exprType == condition.EXPR_TYPE_ENV {
			//environment variable based condition
			envFlagValue := os.Getenv(conditionOperation.LHS)
			logger.Debugf("environment flag = %v, val = %v", conditionOperation.LHS, envFlagValue)
			if envFlagValue != "" {
				conditionOperation.LHS = envFlagValue
				op := conditionOperation.Operator
				exprResult := op.Eval(conditionOperation.LHS, conditionOperation.RHS)
				if exprResult {
					actionID = dispatch.actionID
					handlerCfg = dispatch.handlerCfg
				}
			}
		}

		if actionID != "" {
			logger.Debugf("dispatch resolved with the actionId - %v", actionID)
			break
		}
	}

	//If no dispatch is found, use default action
	if actionID == "" {
		actionID = h.defaultActionID
		handlerCfg = h.defaultHandlerCfg
		logger.Debugf("dispatch not resolved. Continue with default action - %v", actionID)
	}

	return actionID, handlerCfg
}

//Dispatch holds dispatch actionId and condition
type Dispatch struct {
	actionID   string
	condition  string
	handlerCfg *trigger.HandlerConfig
}

