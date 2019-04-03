package jsexec

import (
	"github.com/project-flogo/core/data/coerce"
)

// Settings are the jsexec settings
type Settings struct {
	Script string `md:"script"`
}

// Input is the input into the javascript engine
type Input struct {
	Parameters map[string]interface{} `md:"parameters"`
}

// FromMap converts the values from a map into the struct Input
func (i *Input) FromMap(values map[string]interface{}) error {
	parameters, err := coerce.ToObject(values["parameters"])
	if err != nil {
		return err
	}
	i.Parameters = parameters
	return nil
}

// ToMap converts the struct Input into a map
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"parameters": i.Parameters,
	}
}

// Output is the output from the javascript engine
type Output struct {
	Error        bool                   `md:"error"`
	ErrorMessage string                 `md:"errorMessage"`
	Result       map[string]interface{} `md:"result"`
}

// FromMap converts the values from a map into the struct Output
func (o *Output) FromMap(values map[string]interface{}) error {
	errorValue, err := coerce.ToBool(values["error"])
	if err != nil {
		return err
	}
	o.Error = errorValue
	errorMessage, err := coerce.ToString(values["errorMessage"])
	if err != nil {
		return err
	}
	o.ErrorMessage = errorMessage
	result, err := coerce.ToObject(values["result"])
	if err != nil {
		return err
	}
	o.Result = result
	return nil
}

// ToMap converts the struct Output into a map
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"error":        o.Error,
		"errorMessage": o.ErrorMessage,
		"result":       o.Result,
	}
}
