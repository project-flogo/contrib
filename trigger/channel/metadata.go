package channel

const ovData = "data"

type HandlerSettings struct {
	Channel string `md:"channel,required"`
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
