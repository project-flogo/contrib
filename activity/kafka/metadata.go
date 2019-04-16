package kafka

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	BrokerUrls string `md:"brokerUrls,required"`
	User       string `md:"user"`
	Password   string `md:"password"`
	TrustStore string `md:"truststore"`
}
type Input struct {
	Topic   string `md:"topic,required"`
	Message string `md:"message,required"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{

		"topic":   i.Topic,
		"message": i.Message,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error

	i.Topic, err = coerce.ToString(values["topic"])
	if err != nil {
		return err
	}
	i.Message, err = coerce.ToString(values["message"])
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
