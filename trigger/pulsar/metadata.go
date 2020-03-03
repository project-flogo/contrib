package sample

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	URL                  string `md:"url,required"`
	AthenzAuthentication string `md:"athenzauth"`
	CertFile             string `md:"certFile"`
	KeyFile              string `md:"keyFile"`
}

type HandlerSettings struct {
	Topic        string `md:"topic,required"`
	Subscription string `md:"subscription,required"`
}

type Output struct {
	Message string `md:"message"`
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.Message, err = coerce.ToString(values["message"])
	if err != nil {
		return err
	}

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"message": o.Message,
	}
}
