package kafka

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	BrokerUrls string `md:"brokerurls,required"`
	Topic      string `md:"topic,required"`
	Message    string `md:"message,required"`
	User       string `md:"user"`
	Password   string `md:"password"`
	TrustStore string `md:"truststore"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"brokerurls": i.BrokerUrls,
		"topic":      i.Topic,
		"message":    i.Message,
		"user":       i.User,
		"password":   i.Password,
		"truststore": i.TrustStore,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.BrokerUrls, err = coerce.ToString(values["brokerurls"])
	if err != nil {
		return err
	}
	i.Topic, err = coerce.ToString(values["topic"])
	if err != nil {
		return err
	}
	i.Message, err = coerce.ToString(values["message"])
	if err != nil {
		return err
	}
	i.User, err = coerce.ToString(values["user"])
	if err != nil {
		return err
	}
	i.Password, err = coerce.ToString(values["password"])
	if err != nil {
		return err
	}
	i.TrustStore, err = coerce.ToString(values["truststore"])
	if err != nil {
		return err
	}

	return nil
}

type Output struct {
	Partition int32 `md:"partition"`
	OffSet    int64 `md:"offset"`
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
