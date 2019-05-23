package kafka

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	BrokerUrls string `md:"brokerUrls,required"` // The Kafka cluster to connect to
	User       string `md:"user"`                // If connecting to a SASL enabled port, the user id to use for authentication
	Password   string `md:"password"`            // If connecting to a SASL enabled port, the password to use for authentication
	TrustStore string `md:"trustStore"`          // If connecting to a TLS secured port, the directory containing the certificates representing the trust chain for the connection. This is usually just the CACert used to sign the server's certificate
}
type HandlerSettings struct {
	Topic      string `md:"topic,required"` // The Kafka topic on which to listen for messageS
	Partitions string `md:"partitions"`     // The specific partitions to consume messages from
	Offset     int64  `md:"offset"`         // The offset to use when starting to consume messages, default is set to Newest
}

type Output struct {
	Message string `md:"message"` // The message that was consumed
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"message": o.Message,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.Message, err = coerce.ToString(values["message"])
	if err != nil {
		return err
	}

	return nil
}
