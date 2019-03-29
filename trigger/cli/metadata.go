package cli

import "github.com/project-flogo/core/data/coerce"

const ovArgs = "args"
const ovFlags = "flags"

type Settings struct {
	SingleCmd bool   `md:"singleCmd"`  // Indicates that this CLI runs only one command/handler
	Usage     string `md:"usage"`      // The usage details of the CLI
	Long      string `md:"long"`       // The description of the CLI
}

type HandlerSettings struct {
	FlagDesc []interface{} `md:"flags"`
	Usage    string        `md:"usage"` // The usage details of the command
	Short    string        `md:"short"` // A short description of the command
	Long     string        `md:"long"`  // The description of the command
}

type Output struct {
	Args  []interface{}          `md:"args"`   // An array of the command line arguments
	Flags map[string]interface{} `md:"flags"`  // A map of the command line flags
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
	Data interface{} `md:"data"` // The data that the command outputs
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
