package eftl

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
	"github.com/project-flogo/contrib/trigger/eftl/eftlHelpers"
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
	connection *eftlHelpers.Connection
	stop       chan bool
}

// New implements trigger.Factory.New
func (f *Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	s := &Settings{}
	err := metadata.MapToStruct(config.Settings, s, true)
	if err != nil {
		return nil, err
	}

	return &Trigger{metadata: f.Metadata(), config: config}, nil
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

func (t *Trigger) newActionHandler(handler trigger.Handler) error {
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
	options := &eftlHelpers.Options{
		ClientID:  id.(string),
		Username:  user.(string),
		Password:  password.(string),
		TLSConfig: tlsConfig,
	}

	url := t.config.Settings[settingURL]
	errorsChannel := make(chan error, 1)
	connectVal, err := eftlHelpers.Connect(url.(string), options, errorsChannel)
	if err != nil {
		t.logger.Errorf("connection failed: %s", err)
		return err
	}
	t.connection = connectVal

	messages := make(chan eftlHelpers.Message, 1000)
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
				value := message["content"]
				content, ok := value.([]byte)
				if !ok {
					content = []byte{}
				}
				replyTo := ""
				var js map[string]interface{}
				if json.Unmarshal(content, &js) == nil {
					replyTo = "json"
				} else {
					replyTo = "jsonString"
				}
				out := &Output{}
				out.QueryParams = make(map[string]string)
				out.PathParams = make(map[string]string)
				out.Params = make(map[string]string)
				out.Content = js

				results, err := handler.Handle(context.Background(), out)
				if err != nil {
					t.logger.Errorf("failed to get new handler data: %v", err)
					return
				}

				err = t.connection.Publish(eftlHelpers.Message{
					"_dest":   replyTo,
					"content": results,
				})
			case err := <-errorsChannel:
				t.logger.Errorf("connection error: %s", err)
			case <-t.stop:
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
