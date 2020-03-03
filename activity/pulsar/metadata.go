package sample

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	URL                  string `md:"url,required"`
	Topic                string `md:"topic,required"`
	AthenzAuthentication string `md:"athenzauth"`
	CertFile             string `md:"certFile"`
	KeyFile              string `md:"keyFile"`
}

type Input struct {
	Payload string `md:"payload,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	var err error
	r.Payload, err = coerce.ToString(values["payload"])

	return err
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"payload": r.Payload,
	}
}
