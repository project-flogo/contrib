package sample

import (
	"context"

	"github.com/apache/pulsar/pulsar-client-go/pulsar"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

var triggerMd = trigger.NewMetadata(&Settings{}, &HandlerSettings{}, &Output{})

func init() {
	_ = trigger.Register(&Trigger{}, &Factory{})
}

type Trigger struct {
	client   pulsar.Client
	handlers []*Handler
}
type Handler struct {
	handler  trigger.Handler
	consumer pulsar.Consumer
}

type Factory struct {
}

var logger log.Logger

func (*Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	s := &Settings{}
	err := metadata.MapToStruct(config.Settings, s, true)
	if err != nil {
		return nil, err
	}
	auth := getAuthentication(s)

	clientOps := pulsar.ClientOptions{
		URL:            s.URL,
		Authentication: auth,
	}
	client, err := pulsar.NewClient(clientOps)

	if err != nil {
		return nil, err
	}

	return &Trigger{client: client}, nil
}

func (f *Factory) Metadata() *trigger.Metadata {
	return triggerMd
}

// Metadata implements trigger.Trigger.Metadata
func (t *Trigger) Metadata() *trigger.Metadata {
	return triggerMd
}

func (t *Trigger) Initialize(ctx trigger.InitContext) error {

	logger = ctx.Logger()

	// Init handlers
	for _, handler := range ctx.GetHandlers() {

		s := &HandlerSettings{}
		err := metadata.MapToStruct(handler.Settings(), s, true)
		if err != nil {
			return err
		}
		consumer, err := t.client.Subscribe(pulsar.ConsumerOptions{
			Topic:            s.Topic,
			SubscriptionName: s.Subscription,
		})
		if err != nil {
			return err
		}

		t.handlers = append(t.handlers, &Handler{handler: handler, consumer: consumer})
	}

	return nil
}

// Start implements util.Managed.Start
func (t *Trigger) Start() error {
	for _, handler := range t.handlers {
		go consume(handler)
	}
	return nil
}

// Stop implements util.Managed.Stop
func (t *Trigger) Stop() error {
	for _, handler := range t.handlers {
		handler.consumer.Close()
	}
	return nil
}

func getAuthentication(s *Settings) pulsar.Authentication {
	if s.AthenzAuthentication != "" {
		return pulsar.NewAuthenticationAthenz(s.AthenzAuthentication)

	}
	if s.CertFile != "" && s.KeyFile != "" {
		return pulsar.NewAuthenticationTLS(s.CertFile, s.KeyFile)
	}
	return nil
}

func consume(handler *Handler) {

	for {

		msg, err := handler.consumer.Receive(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		out := &Output{}
		out.Message = string(msg.Payload())
		logger.Debugf("Message recieved [%v]", out.Message)
		// Do something with the message
		_, err = handler.handler.Handle(context.Background(), out)

		if err == nil {
			// Message processed successfully
			handler.consumer.Ack(msg)
		} else {
			// Failed to process messages
			handler.consumer.Nack(msg)
		}

	}
}
