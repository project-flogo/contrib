package loadtester

const ovData = "data"

type Settings struct {
	Concurrency int         `md:"concurrency"`
	Duration    int         `md:"duration"`
	Data        interface{} `md:"data"`
	Handler     string      `md:"handler"`
	StartDelay  int         `md:"startDelay"`
}

type Output struct {
	Data interface{} `md:"data"`
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		ovData: o.Data,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {
	o.Data = values[ovData]
	return nil
}
