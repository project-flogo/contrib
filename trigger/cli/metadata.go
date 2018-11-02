package cli

import "github.com/project-flogo/core/data/coerce"

const ovArgs = "args"

type HandlerSettings struct {
	Command string `md:"command"`
	Default bool `md:"default"`
}

type Output struct {
	Args []interface{} `md:"args"`
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		ovArgs: o.Args,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {
	var err error
	o.Args, err = coerce.ToArray(values[ovArgs])
	return err
}

type Reply struct {
	Data interface{} `md:"data"`
}

func (r *Reply) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"data": r.Data,
	}
}

func (r *Reply) FromMap(values map[string]interface{}) error {
	var err error
	r.Data, err = values["data"]
	return err
}