package cli

import "github.com/project-flogo/core/data/coerce"

const ovArgs = "args"
const ovFlags = "flags"

type Settings struct {
	SingleCmd bool   `md:"singleCmd"`
	Use       string `md:"use"`
	Long      string `md:"long"`
}

type HandlerSettings struct {
	FlagDesc []interface{} `md:"flags"`
	Use      string        `md:"use"`
	Short    string        `md:"short"`
	Long     string        `md:"long"`
}

type Output struct {
	Args  []interface{}          `md:"args"`
	Flags map[string]interface{} `md:"flags"`
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		ovArgs:  o.Args,
		ovFlags: o.Flags,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {
	var err error
	o.Args, err = coerce.ToArray(values[ovArgs])
	o.Flags, err = coerce.ToObject(values[ovFlags])
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
	r.Data, _ = values["data"]
	return nil
}
