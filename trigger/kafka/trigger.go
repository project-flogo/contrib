package kafka

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

var triggerMd = trigger.NewMetadata(&Settings{}, &HandlerSettings{}, &Output{})

func init() {
	_ = trigger.Register(&KafkaTrigger{}, &Factory{})
}

type Factory struct {
}

// Metadata implements trigger.Factory.Metadata
func (*Factory) Metadata() *trigger.Metadata {
	return triggerMd
}

// New implements trigger.Factory.New
func (*Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	s := &Settings{}
	err := metadata.MapToStruct(config.Settings, s, true)
	if err != nil {
		return nil, err
	}

	return &KafkaTrigger{settings: s}, nil
}

type KafkaTrigger struct {
	settings      *Settings
	conn          *KafkaConnection
	kafkaHandlers []*KafkaHandler
}

func (t *KafkaTrigger) Initialize(ctx trigger.InitContext) error {

	var err error
	t.conn, err = getKafkaConnection(ctx.Logger(), t.settings)

	for _, handler := range ctx.GetHandlers() {
		kafkaHandler, err := NewKafkaHandler(ctx.Logger(), handler, t.conn.Connection())
		if err != nil {
			return err
		}
		t.kafkaHandlers = append(t.kafkaHandlers, kafkaHandler)
	}

	return err
}

func (t *KafkaTrigger) Start() error {

	for _, handler := range t.kafkaHandlers {
		_ = handler.Start()
	}

	return nil
}

// Stop implements ext.Trigger.Stop
func (t *KafkaTrigger) Stop() error {

	for _, handler := range t.kafkaHandlers {
		_ = handler.Stop()
	}

	_ = t.conn.Stop()
	return nil
}

func NewKafkaHandler(logger log.Logger, handler trigger.Handler, consumer sarama.Consumer) (*KafkaHandler, error) {

	kafkaHandler := &KafkaHandler{logger: logger, shutdown:make(chan struct{})}

	handlerSetting := &HandlerSettings{}
	err := metadata.MapToStruct(handler.Settings(), handlerSetting, true)
	if err != nil {
		return nil, err
	}

	if handlerSetting.Topic == "" {
		return nil, fmt.Errorf("topic string was not provided for handler: [%s]", handler)
	}

	logger.Debugf("Subscribing to topic [%s]", handlerSetting.Topic)

	offset := sarama.OffsetNewest

	//offset
	if handlerSetting.Offset != 0 {
		offset = handlerSetting.Offset
	}

	var partitions []int32

	validPartitions, err := consumer.Partitions(handlerSetting.Topic)
	logger.Debugf("Valid partitions for topic [%s] detected as: [%v]", handlerSetting.Topic, validPartitions)

	if handlerSetting.Partitions != "" {
		parts := strings.Split(handlerSetting.Partitions, ",")
		for _, p := range parts {
			n, err := strconv.Atoi(p)
			if err == nil {
				for _, validPartition := range validPartitions {
					if int32(n) == validPartition {
						partitions = append(partitions, int32(n))
						break
					}
					logger.Errorf("Configured partition [%d] on topic [%s] does not exist and will not be subscribed", n, handlerSetting.Topic)
				}
			} else {
				logger.Warnf("Partition [%s] specified for handler [%s] is not a valid number and was discarded", p, handler)
			}
		}
	} else {
		partitions = validPartitions
	}

	for _, partition := range partitions {
		logger.Debugf("Creating PartitionConsumer for partition: [%s:%d]", handlerSetting.Topic, partition)
		partitionConsumer, err := consumer.ConsumePartition(handlerSetting.Topic, partition, offset)
		if err != nil {
			logger.Errorf("Creating PartitionConsumer for valid partition: [%s:%d] failed for reason: %s", handlerSetting.Topic, partition, err)
			return nil, err
		}
		kafkaHandler.consumers = append(kafkaHandler.consumers, partitionConsumer)
	}

	return kafkaHandler, nil
}

type KafkaHandler struct {
	shutdown  chan struct{}
	logger    log.Logger
	handler   trigger.Handler
	consumers []sarama.PartitionConsumer
}

func (h *KafkaHandler) consumePartition(consumer sarama.PartitionConsumer) {
	for {
		select {
		case err := <-consumer.Errors():
			if err == nil {
				//was shutdown
				return
			}
			time.Sleep(time.Millisecond * 100)
		case <- h.shutdown:
			return
		case msg := <-consumer.Messages():

			if h.logger.DebugEnabled() {
				h.logger.Debugf("Kafka subscriber triggering action from topic [%s] on partition [%d] with key [%s] at offset [%d]",
					msg.Topic, msg.Partition, msg.Key, msg.Offset)

				h.logger.Debugf("Kafka message: '%s'", string(msg.Value))
			}

			out := &Output{}
			out.Message = string(msg.Value)

			_, err := h.handler.Handle(context.Background(), out)
			if err != nil {
				h.logger.Errorf("Run action for handler [%s] failed for reason [%s] message lost", h.handler.Name(), err)
			}
		}
	}
}

func (h *KafkaHandler) Start() error {

	for _, consumer := range h.consumers {
		h.consumePartition(consumer)
	}

	return nil
}

func (h *KafkaHandler) Stop() error {

	close(h.shutdown)

	for _, consumer := range h.consumers {
		_ = consumer.Close()
	}
	return nil
}
