package loadtester

const ovData = "data"

type Settings struct {
	Concurrency int         `md:"concurrency"`  // The level of concurrency, default: 5
	Duration    int         `md:"duration"`     // The duration of the test in seconds, default: 60
	Data        interface{} `md:"data"`         // Optional data to pass along to the action
	Handler     string      `md:"handler"`      // The named handler to test, defaults to the first handler
	StartDelay  int         `md:"startDelay"`   // The start delay of the test in seconds, default: 30
}

type Output struct {
	Data interface{} `md:"data"`  // The data from the settings to pass along
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
