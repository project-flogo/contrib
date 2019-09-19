package runaction

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Ref    string `md:"ref"` //The ref to action
	ResURI string `md:"resURI"`
}

type Input struct {
	Input interface{} `md:"input"`
}

type Output struct {
	Output map[string]interface{} `md:"output"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"input": i.Input,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	i.Input = values["input"]

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"output": o.Output,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error

	o.Output, err = coerce.ToObject(values["output"])
	if err != nil {
		return err
	}

	return nil
}
