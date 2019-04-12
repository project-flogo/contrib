package kafka

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
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
	_ = trigger.Register(&KafkaSubTrigger{}, &Factory{})
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

	return &KafkaSubTrigger{settings: s}, nil
}

type _topichandler struct {
	topic      string
	offset     int64
	group      string
	partitions []int32
}

type _kafkaParms struct {
	brokers  []string
	handlers []_topichandler
}

type KafkaSubTrigger struct {
	settings           *Settings
	handlers           []trigger.Handler
	kafkaParms         _kafkaParms
	shutdownChan       *chan struct{}
	signals            *chan os.Signal
	kafkaConfig        *sarama.Config
	kafkaConsumer      *sarama.Consumer
	partitionConsumers *map[string]sarama.PartitionConsumer
	logger             log.Logger
}

func (t *KafkaSubTrigger) Initialize(ctx trigger.InitContext) error {
	t.handlers = ctx.GetHandlers()
	t.logger = ctx.Logger()
	err := initKafkaParams(t)
	return err
}

func (t *KafkaSubTrigger) Start() error {
	shutdownChan := make(chan struct{})
	t.shutdownChan = &shutdownChan
	signals := make(chan os.Signal, 1)
	t.signals = &signals
	signal.Notify(*t.signals, os.Interrupt)
	err := run(t)
	//log.Debug("KafkaSubTrigger Started")
	return err
}

// Stop implements ext.Trigger.Stop
func (t *KafkaSubTrigger) Stop() error {
	//unsubscribe from topic
	if t.partitionConsumers == nil {
		t.logger.Debug("Closed called for a subscriber with no running consumers")
		t.logger.Debug("Stopped")
		return nil
	}
	for id, partitionConsumer := range *t.partitionConsumers {
		partitionConsumer.Close()
		t.logger.Debug("Closed partition consumer:", id)
	}
	if t.kafkaConsumer != nil {
		(*t.kafkaConsumer).Close()
		t.logger.Debug("Closed kafka consumer")
	}
	t.logger.Debug("Stopped")
	return nil
}

func run(t *KafkaSubTrigger) error {
	kafkaConsumer, err := sarama.NewConsumer(t.kafkaParms.brokers, t.kafkaConfig)
	if err != nil {
		return fmt.Errorf("failed to create Kafka consumer for reason [%s]", err)
	}
	t.kafkaConsumer = &kafkaConsumer
	consumers := make(map[string]sarama.PartitionConsumer)
	t.partitionConsumers = &consumers

	for id, handler := range t.kafkaParms.handlers {
		validPartitions, err := kafkaConsumer.Partitions(handler.topic)
		if err != nil {
			return fmt.Errorf("failed to get valid partitions for topic [%s] for reason [%s].  Aborting subscriber",
				handler.topic, err)
		}
		t.logger.Debugf("Subscribing to topic [%s]", handler.topic)

		t.logger.Debugf("Valid partitions for topic [%s] detected as: [%v]", handler.topic, validPartitions)
		if handler.partitions == nil { //subscribe to all valid partitions
			handler.partitions = validPartitions
			for _, part := range validPartitions {
				startConsumer(t, part, id)
			}
		} else { //subscribe to a subset of valid partitions
			configPartitions := handler.partitions
			for _, confPart := range configPartitions {
				for _, valPart := range validPartitions {
					if confPart == valPart {
						startConsumer(t, confPart, id)
						break
					}
					t.logger.Errorf("Configured partition [%d] on topic [%s] does not exist and will not be subscribed", confPart, handler.topic)
				}
			}
		}
		if len(*t.partitionConsumers) < 1 {
			return fmt.Errorf("Kafka consumer is not configured for any valid partitions")
		}
		t.logger.Debugf("Kafka consumers for topic [%s] started", handler.topic)
	}
	return nil
}

func startConsumer(t *KafkaSubTrigger, part int32, id int) error {
	t.logger.Debugf("Creating PartitionConsumer for valid partition: [%s:%d]", t.kafkaParms.handlers[id].topic, part)
	consumer := *t.kafkaConsumer
	partitionConsumer, err := consumer.ConsumePartition(t.kafkaParms.handlers[id].topic, part, t.kafkaParms.handlers[id].offset)
	if err != nil {
		t.logger.Errorf("Creating PartitionConsumer for valid partition: [%s:%d] failed for reason: %s", t.kafkaParms.handlers[id].topic, part, err)
		return err
	}
	consumerName := fmt.Sprintf("%d_%d", part, id)
	(*t.partitionConsumers)[consumerName] = partitionConsumer
	go consumePartition(t, partitionConsumer, part)
	return nil
}

func consumePartition(t *KafkaSubTrigger, consumer sarama.PartitionConsumer, part int32) {
	for {
		select {
		case err := <-consumer.Errors():
			if err == nil {
				//been shutdown
				return
			}
			t.logger.Warnf("PartitionConsumer [%d] got error: [%s]", part, err)
			time.Sleep(time.Millisecond * 300)
		case msg := <-consumer.Messages():
			onMessage(t, msg)
		case <-*t.signals:
			t.logger.Infof("Partition consumer got SIGINT; exiting")
			*t.shutdownChan <- struct{}{}
			return
		case <-*t.shutdownChan:
			t.logger.Infof("Partition consumer got SHUTDOWN signal; exiting")
			return
		}
	}
}

/*
func getTopics(t *KafkaSubTrigger) []string {
	return strings.Split(t.settings.Topics, ",")
}
*/

func initKafkaParams(t *KafkaSubTrigger) error {
	brokersString := t.settings.BrokerUrls
	if brokersString == "" {
		return fmt.Errorf("BrokerUrl not provided")
	}
	brokers := strings.Split(brokersString, ",")
	if len(brokers) < 1 {
		return fmt.Errorf("BrokerUrl [%s] is invalid, require at least one broker", brokersString)
	}
	t.kafkaParms.brokers = make([]string, len(brokers))
	for brokerNo, broker := range brokers {
		err := validateBrokerUrl(broker)
		if err != nil {
			return fmt.Errorf("BrokerUrl [%s] format invalid for reason: [%s]", broker, err.Error())
		}
		t.kafkaParms.brokers[brokerNo] = broker
	}
	//clientKeystore
	/*
		Its worth mentioning here that when the keystore for kafka is created it must support RSA keys via
		the -keyalg RSA option.  If not then there will be ZERO overlap in supported cipher suites with java.
		see:   https://issues.apache.org/jira/browse/KAFKA-3647
		for more info
	*/
	if trustStore := t.settings.TrustStore; len(trustStore) > 0 {
		trustPool, err := getCerts(trustStore)
		if err != nil {
			return err
		}
		config := tls.Config{
			RootCAs:            trustPool,
			InsecureSkipVerify: true}
		t.kafkaConfig.Net.TLS.Enable = true
		t.kafkaConfig.Net.TLS.Config = &config
	}
	// SASL
	if t.settings.User != "" {
		var password string
		user := t.settings.User
		if len(user) > 0 {
			if len(t.settings.Password) < 1 {
				return fmt.Errorf("password not provided for user: %s", user)
			}
			password = t.settings.Password
			t.kafkaConfig.Net.SASL.Enable = true
			t.kafkaConfig.Net.SASL.User = user
			t.kafkaConfig.Net.SASL.Password = password
		}
	}

	// _topichandlers section
	if len(t.handlers) == 0 {
		return fmt.Errorf("Kafka trigger requires at least one handler containing a valid topic name")
	}

	t.kafkaParms.handlers = make([]_topichandler, len(t.handlers))

	for handlerNum, handler := range t.handlers {
		handlerSetting := &HandlerSettings{}
		err := metadata.MapToStruct(handler.Settings(), handlerSetting, true)
		if err != nil {
			return err
		}

		if handlerSetting.Topic == "" {
			return fmt.Errorf("topic string was not provided for handler: [%s]", handler)
		}
		t.kafkaParms.handlers[handlerNum].topic = handlerSetting.Topic

		//offset
		if handlerSetting.OffSet != 0 {
			i := handlerSetting.OffSet
			if err != nil {
				t.logger.Warnf("Offset [%s] specified for handler [%s] is not a valid number, using latest for offset",
					handlerSetting.OffSet, handler)
				t.kafkaParms.handlers[handlerNum].offset = sarama.OffsetNewest
			} else {
				t.kafkaParms.handlers[handlerNum].offset = int64(i)
			}
		} else {
			t.kafkaParms.handlers[handlerNum].offset = sarama.OffsetNewest
		}

		if handlerSetting.Partition != "" {
			partitions := handlerSetting.Partition
			i := 0
			parts := strings.Split(partitions, ",")
			t.kafkaParms.handlers[handlerNum].partitions = make([]int32, len(partitions))
			for _, p := range parts {
				n, err := strconv.Atoi(p)
				if err == nil {
					t.kafkaParms.handlers[handlerNum].partitions[i] = int32(n)
					i++
				} else {
					fmt.Errorf("Partition [%s] specified for handler [%s] is not a valid number and is discarded",
						p, handler)
				}
			}
		}

		//group
		if handlerSetting.Group != "" {
			t.kafkaParms.handlers[handlerNum].group = handlerSetting.Group
		}
	}
	return nil
}

func getCerts(trustStore string) (*x509.CertPool, error) {
	certPool := x509.NewCertPool()
	fileInfo, err := os.Stat(trustStore)
	if err != nil {
		return certPool, fmt.Errorf("Truststore [%s] does not exist", trustStore)
	}
	switch mode := fileInfo.Mode(); {
	case mode.IsDir():
		break
	case mode.IsRegular():
		return certPool, fmt.Errorf("Truststore [%s] is not a directory.  Must be a directory containing trusted certificates in PEM format",
			trustStore)
	}
	trustedCertFiles, err := ioutil.ReadDir(trustStore)
	if err != nil || len(trustedCertFiles) == 0 {
		return certPool, fmt.Errorf("Failed to read trusted certificates from [%s]  Must be a directory containing trusted certificates in PEM format", trustStore)
	}
	for _, trustCertFile := range trustedCertFiles {
		fqfName := fmt.Sprintf("%s%c%s", trustStore, os.PathSeparator, trustCertFile.Name())
		trustCertBytes, err := ioutil.ReadFile(fqfName)
		if err != nil {
			fmt.Errorf("Failed to read trusted certificate [%s] ... continueing", trustCertFile.Name())
		}
		certPool.AppendCertsFromPEM(trustCertBytes)
	}
	if len(certPool.Subjects()) < 1 {
		return certPool, fmt.Errorf("Failed to read trusted certificates from [%s]  After processing all files in the directory no valid trusted certs were found", trustStore)
	}
	return certPool, nil
}

//Ensure that this string meets the host:port definition of a kafka hostspec
//Kafka calls it a url but its really just host:port, which for numeric ip addresses is not a valid URI
//technically speaking.
func validateBrokerUrl(broker string) error {
	hostport := strings.Split(broker, ":")
	if len(hostport) != 2 {
		return fmt.Errorf("BrokerUrl must be composed of sections like \"host:port\"")
	}
	i, err := strconv.Atoi(hostport[1])
	if err != nil || i < 0 || i > 32767 {
		return fmt.Errorf("Port specification [%s] is not numeric and between 0 and 32767", hostport[1])
	}
	return nil
}

func onMessage(t *KafkaSubTrigger, msg *sarama.ConsumerMessage) {
	if msg == nil {
		return
	}

	t.logger.Debugf("Kafka subscriber triggering job from topic [%s] on partition [%d] with key [%s] at offset [%d]",
		msg.Topic, msg.Partition, msg.Key, msg.Offset)

	for _, handler := range t.handlers {

		//actionID := action.Get(handler.ActionId)
		//log.Debugf("Found action: '%+x' for ActionID: %s", actionID, handler.ActionId
		out := &Output{}

		out.Message = string(msg.Value)
		fmt.Println("Output is ...",out.Message)
		//if(t.metadata.Metadata.OutPuts

		_, err := handler.Handle(context.Background(), out)

		if err != nil {
			t.logger.Errorf("Run action for handler [%s] failed for reason [%s] message lost", handler, err)
		}
	}

}
