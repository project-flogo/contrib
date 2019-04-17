package kafka

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	BrokerUrls string `md:"brokerUrls,required"` // The Kafka cluster to connect to
	User       string `md:"user"`       // If connecting to a SASL enabled port, the user id to use for authentication
	Password   string `md:"password"`   // If connecting to a SASL enabled port, the password to use for authentication
	TrustStore string `md:"trustStore"` // If connecting to a TLS secured port, the directory containing the certificates representing the trust chain for the connection. This is usually just the CACert used to sign the server's certificate
	Topic      string `md:"topic,required"` // The Kafka topic on which to place the message
}
type Input struct {
	Message string `md:"message,required"` // The message to send
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"message": i.Message,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Message, err = coerce.ToString(values["message"])
	return err
}

type Output struct {
	Partition int32 `md:"partition"` // Documents the partition that the message was placed on
	OffSet    int64 `md:"offset"`    // Documents the offset for the message
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"partition": o.Partition,
		"offset":    o.OffSet,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.Partition, err = coerce.ToInt32(values["partition"])
	if err != nil {
		return err
	}

	o.OffSet, err = coerce.ToInt64(values["offset"])
	if err != nil {
		return err
	}

	return nil
}
