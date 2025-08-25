package json2xml

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	Json       map[string]interface{} `md:"json"`
	XmlRootTag string                 `md:"xmlRootTag"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"json":       i.Json,
		"xmlRootTag": i.XmlRootTag,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error
	i.Json, err = coerce.ToObject(values["json"])
	i.XmlRootTag, err = coerce.ToString(values["xmlRootTag"])
	if err != nil {
		return err
	}
	return nil
}

type Output struct {
	Xml []byte `md:"xml"`
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"xml": o.Xml,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {
	var err error
	o.Xml, err = coerce.ToBytes(values["xml"])
	if err != nil {
		return err
	}

	return nil
}
