package kafkasub

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	BrokerUrls string `md:"brokerurls,required"`
	User       string `md:"user"`
	Password   string `md:"password"`
	TrustStore string `md:"truststore"`
}
type HandlerSettings struct {
	Topic     string `md:"topic,required"`
	Partition string `md:"partition"`
	Group     string `md:"group"`
	OffSet    int64  `md:"offset"`
}

type Output struct {
	Message string `md:"string"`
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
