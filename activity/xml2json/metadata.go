package xml2json

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	XmlData interface{} `md:"xmlData"` //
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"xmlData": i.XmlData,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error
	i.XmlData, err = coerce.ToAny(values["xmlData"])
	if err != nil {
		return err
	}
	return nil
}

type Output struct {
	JsonObject map[string]interface{} `md:"jsonObject"` // The HTTP response data
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"jsonObject": o.JsonObject,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.JsonObject, err = coerce.ToObject(values["jsonObject"])
	if err != nil {
		return err
	}

	return nil
}
