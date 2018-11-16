package eftl

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	//"os"

	"github.com/project-flogo/core/trigger"
	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/data/metadata"
	"github.com/mashling/commons/lib/eftl"
	"github.com/mashling/commons/lib/util"
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
	logger     log.Logger
	//handlers   map[string]*OptimizedHandler
	connection *eftl.Connection
	stop       chan bool
}

// New implements trigger.Factory.New
func (f *Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	s := &Settings{}
	err := metadata.MapToStruct(config.Settings, s, true)
	if err != nil {
		return nil, err
	}

	return &Trigger{metadata: f.Metadata(),config: config}, nil
}

// Init implements trigger.Init
func (t *Trigger) Initialize(ctx trigger.InitContext) error {
	for _, handler := range ctx.GetHandlers() {

		s := &HandlerSettings{}
		err := metadata.MapToStruct(handler.Settings(), s, true)
		if err != nil {
			return err
		}
		err = t.newActionHandler(handler)
		if err != nil {
			return err
		}

	}
	return nil
}

func (t *Trigger) newActionHandler(handler trigger.Handler) error{
	fmt.Println("Inside Trigger action handler")
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	ca := t.config.Settings[settingCA]
	if ca != "" {
		certificate, err := ioutil.ReadFile(ca.(string))
		if err != nil {
			t.logger.Errorf("can't open certificate", err)
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
	fmt.Println("ID : ", id)
	options := &eftl.Options{
		ClientID:  id.(string),
		Username:  user.(string),
		Password:  password.(string),
		TLSConfig: tlsConfig,
	}

	url := t.config.Settings[settingURL]
	fmt.Println("URL : ", url)
	errorsChannel := make(chan error, 1)
	connectVal, err := eftl.Connect(url.(string), options, errorsChannel)
	if err != nil {
		t.logger.Errorf("connection failed: %s", err)
		return err
	}
	t.connection = connectVal
	messages := make(chan eftl.Message, 1000)
	dest := handler.Settings()
	matcher := fmt.Sprintf("{\"_dest\":\"%s\"}", dest[settingDest])
	_, err = t.connection.Subscribe(matcher, "", messages)
	if err != nil {
		t.logger.Errorf("subscription failed: %s", err)
		return err
	}

	t.stop = make(chan bool, 1)
	go func() {
		for {
			select {
			case message := <-messages:
				fmt.Println("Inside case")
				value := message["_dest"]
				dest, ok := value.(string)
				if !ok {
					t.logger.Errorf("dest is required for valid message")
					continue
				}

				value = message["content"]
				fmt.Println("value :", value)
				content, ok := value.([]byte)
				if !ok {
					content = []byte{}
				}
				fmt.Println("Dest val:", dest)
				fmt.Println("Content val:", content)
				err = t.RunAction(content,handler)
				if err != nil{
					t.logger.Errorf(" RunAction failed: %s", err)
				}
			case err := <-errorsChannel:
				t.logger.Errorf("connection error: %s", err)
			case <-t.stop:
				fmt.Println("inside stop")
				return
			}
		}
	}()
	return nil
}


// Start implements ext.Trigger.Start
func (t *Trigger) Start() error {
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
func (t *Trigger) RunAction(content []byte, handler trigger.Handler) error {
	fmt.Println("Inside Runaction")
	fmt.Println("content :", string(content))

	replyTo, data := t.constructStartRequest(content)
	fmt.Println("data :", data)
	fmt.Println("replyto :", replyTo)

	if replyTo == "" {
		t.logger.Errorf("reply data is empty")
		return nil
	}

	replyData, err := handler.Handle(context.Background(), data)
	if err != nil {
		t.logger.Errorf("Error starting action: %v", err)
		return err
	}

	reply, err := util.Marshal(replyData)
	if err != nil {
		t.logger.Errorf("failed to marshal reply data: %v", err)
		return err
	}
	fmt.Println("replyTo :", replyTo)
	fmt.Println("reply :", reply)
	err = t.connection.Publish(eftl.Message{
		"_dest":   replyTo,
		"content": reply,
	})
	if err != nil {
		t.logger.Errorf("failed to send reply data: %v", err)
	}
	return nil
}

func (t *Trigger) constructStartRequest(message []byte) (string, *Output) {

	var content map[string]interface{}
	err := util.Unmarshal("", message, &content)
	fmt.Println("content val:",content)
	if err != nil {
		t.logger.Errorf("Error unmarshaling message %s", err.Error())
	}

	replyTo := ""
	pathParams := make(map[string]string)
	queryParams := make(map[string]string)

	mime := ""
	if value, ok := content[util.MetaMIME].(string); ok {
		mime = value
	}
	fmt.Println("Value of mime :", mime)
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

	out := &Output{}
	out.PathParams = pathParams
	out.Params = pathParams
	out.QueryParams = queryParams
	out.Content = content

	return replyTo, out
}
