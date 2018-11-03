package cli

import "github.com/project-flogo/core/data/coerce"

const ovArgs = "args"

type Settings struct {
	SingleCmd  bool   `md:"singleCmd"`
	DefaultCmd string `md:"defaultCmd"`
	Use        string `md:"use"`   //"flogo [flags] [command]",
	Short      string `md:"short"` //"flogo cli",
	Long       string `md:"long"`  //`flogo command line interface for flogo applications`,
}

type HandlerSettings struct {
	FlagDesc   []interface{} `md:"flags"`
	Use        string `md:"use"`   //"flogo [flags] [command]",
	Short      string `md:"short"` //"flogo cli",
	Long       string `md:"long"`  //`flogo command line interface for flogo applications`,
}

type Output struct {
	Args    []interface{} `md:"args"`
	Flags   map[string]interface{} `md:"flags"`
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
	r.Data, _ = values["data"]
	return nil
}
