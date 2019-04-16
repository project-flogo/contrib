package kafka

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = activity.Register(&KafkaActivity{}, New)
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

// MyActivity is a stub for your Activity implementation
type KafkaActivity struct {
	conn  *KafkaConnection
	topic string
}

func New(ctx activity.InitContext) (activity.Activity, error) {
	settings := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), settings, true)
	if err != nil {
		return nil, err
	}

	conn, err := getKafkaConnection(ctx.Logger(), settings)
	if err != nil {
		//ctx.Logger().Errorf("Kafka parameters initialization got error: [%s]", err.Error())
		return nil, err
	}

	act := &KafkaActivity{conn: conn, topic: settings.Topic}
	return act, nil
}

func (act *KafkaActivity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements activity.Activity.Eval
func (act *KafkaActivity) Eval(ctx activity.Context) (done bool, err error) {
	input := &Input{}

	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	if input.Message == "" {
		return false, fmt.Errorf("no message to publish")
	}

	ctx.Logger().Debugf("sending Kafka message")

	msg := &sarama.ProducerMessage{
		Topic: act.topic,
		Value: sarama.StringEncoder(input.Message),
	}

	partition, offset, err := act.conn.Connection().SendMessage(msg)
	if err != nil {
		return false, fmt.Errorf("failed to send Kakfa message for reason [%s]", err.Error())
	}

	output := &Output{}
	output.Partition = partition
	output.OffSet = offset

	if ctx.Logger().DebugEnabled() {
		ctx.Logger().Debugf("Kafka message [%v] sent successfully on partition [%d] and offset [%d]",
			input.Message, partition, offset)
	}

	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}

func getKafkaConnection(logger log.Logger, settings *Settings) (*KafkaConnection, error) {

	connKey := getConnectionKey(settings)

	if conn, ok := connections[connKey]; ok {
		logger.Debugf("Reusing cached Kafka connection [%s]", connKey)
		return conn, nil
	}

	newConn := &KafkaConnection{}

	newConn.kafkaConfig = sarama.NewConfig()
	newConn.kafkaConfig.Producer.Return.Errors = true
	newConn.kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	newConn.kafkaConfig.Producer.Retry.Max = 5
	newConn.kafkaConfig.Producer.Return.Successes = true
	brokerUrls := strings.Split(settings.BrokerUrls, ",")
	brokers := make([]string, len(brokerUrls))

	for brokerNo, broker := range brokerUrls {
		err := validateBrokerUrl(broker)
		if err != nil {
			return nil, fmt.Errorf("BrokerUrl [%s] format invalid for reason: [%v]", broker, err)
		}
		brokers[brokerNo] = broker
	}

	newConn.brokers = brokers
	logger.Debugf("Kafka brokers: [%v]", brokers)

	//clientKeystore
	/*
		Its worth mentioning here that when the keystore for kafka is created it must support RSA keys via
		the -keyalg RSA option.  If not then there will be ZERO overlap in supported cipher suites with java.
		see: https://issues.apache.org/jira/browse/KAFKA-3647
		for more info
	*/
	if trustStore := settings.TrustStore; trustStore != "" {
		if trustPool, err := getCerts(logger, trustStore); err == nil {
			config := tls.Config{
				RootCAs:            trustPool,
				InsecureSkipVerify: true}
			newConn.kafkaConfig.Net.TLS.Enable = true
			newConn.kafkaConfig.Net.TLS.Config = &config

			logger.Debugf("Kafka initialized truststore from [%v]", trustStore)
		} else {
			return nil, err
		}
	}

	// SASL
	if user := settings.User; user != "" {
		password := settings.Password
		newConn.kafkaConfig.Net.SASL.Enable = true
		newConn.kafkaConfig.Net.SASL.User = user
		newConn.kafkaConfig.Net.SASL.Password = password
		logger.Debugf("Kafka SASL params initialized; user [%v]  password[########]", user)
	}

	syncProducer, err := sarama.NewSyncProducer(newConn.brokers, newConn.kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create a Kafka SyncProducer.  Check any TLS or SASL parameters carefully.  Reason given: [%s]", err)
	}
	newConn.syncProducer = syncProducer
	connections[connKey] = newConn
	logger.Debugf("Caching Kafka connection [%s]", connKey)

	return newConn, nil
}

// validateBrokerUrl ensures that this string meets the host:port definition of a kafka host spec
// Kafka calls it a url but its really just host:port, which for numeric ip addresses is not a valid URI
// technically speaking.
func validateBrokerUrl(broker string) error {
	hostPort := strings.Split(broker, ":")
	if len(hostPort) != 2 {
		return fmt.Errorf("BrokerUrl must be composed of sections like \"host:port\"")
	}
	i, err := strconv.Atoi(hostPort[1])
	if err != nil || i < 0 || i > 32767 {
		return fmt.Errorf("port specification [%s] is not numeric and between 0 and 32767", hostPort[1])
	}
	return nil
}

func getCerts(logger log.Logger, trustStore string) (*x509.CertPool, error) {
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
		return certPool, fmt.Errorf("failed to read trusted certificates from [%s]  Must be a directory containing trusted certificates in PEM format", trustStore)
	}

	for _, trustCertFile := range trustedCertFiles {
		fqfName := fmt.Sprintf("%s%c%s", trustStore, os.PathSeparator, trustCertFile.Name())
		trustCertBytes, err := ioutil.ReadFile(fqfName)
		if err != nil {
			logger.Warnf("Failed to read trusted certificate [%s] ... continuing", trustCertFile.Name())
		} else if trustCertBytes != nil {
			certPool.AppendCertsFromPEM(trustCertBytes)
		}
	}

	if len(certPool.Subjects()) < 1 {
		return certPool, fmt.Errorf("failed to read trusted certificates from [%s]  After processing all files in the directory no valid trusted certs were found", trustStore)
	}

	return certPool, nil
}

//////////////////////////////////////////////////
// Connection Support

// todo core should add support for shared connections and replace this
var connections = make(map[string]*KafkaConnection)

type KafkaConnection struct {
	kafkaConfig  *sarama.Config
	brokers      []string
	syncProducer sarama.SyncProducer
}

func (c *KafkaConnection) Connection() sarama.SyncProducer {
	return c.syncProducer
}

func getConnectionKey(settings *Settings) string {

	var connKey string

	connKey += settings.BrokerUrls
	if settings.TrustStore != "" {
		connKey += settings.TrustStore
	}
	if settings.User != "" {
		connKey += settings.User
	}

	return connKey
}


