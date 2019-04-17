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
	"github.com/project-flogo/core/support/log"
)

type KafkaConnection struct {
	kafkaConfig  *sarama.Config
	brokers      []string
	consumer sarama.Consumer
}

func (c *KafkaConnection) Connection() sarama.Consumer {
	return c.consumer
}

func (c *KafkaConnection) Stop() error {
	return c.consumer.Close()
}

func getKafkaConnection(logger log.Logger, settings *Settings) (*KafkaConnection, error) {

	newConn := &KafkaConnection{}

	newConn.kafkaConfig = sarama.NewConfig()
	brokerUrls := strings.Split(settings.BrokerUrls, ",")

	if len(brokerUrls) < 1 {
		return nil, fmt.Errorf("BrokerUrl [%s] is invalid, require at least one broker", settings.BrokerUrls)
	}

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
	if settings.TrustStore != "" {
		if trustPool, err := getCerts(logger, settings.TrustStore); err == nil {
			config := tls.Config{
				RootCAs:            trustPool,
				InsecureSkipVerify: true}
			newConn.kafkaConfig.Net.TLS.Enable = true
			newConn.kafkaConfig.Net.TLS.Config = &config

			logger.Debugf("Kafka initialized truststore from [%v]", settings.TrustStore)
		} else {
			return nil, err
		}
	}

	// SASL
	if settings.User != "" {
		if len(settings.Password) == 0 {
			return nil, fmt.Errorf("password not provided for user: %s", settings.User)
		}
		newConn.kafkaConfig.Net.SASL.Enable = true
		newConn.kafkaConfig.Net.SASL.User = settings.User
		newConn.kafkaConfig.Net.SASL.Password = settings.Password
		logger.Debugf("Kafka SASL params initialized; user [%v]", settings.User)
	}

	kafkaConsumer, err := sarama.NewConsumer(brokers, newConn.kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka consumer for reason [%s]", err)
	}

	newConn.consumer = kafkaConsumer

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
		return certPool, fmt.Errorf("TrustStore [%s] is not a directory.  Must be a directory containing trusted certificates in PEM format",
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
