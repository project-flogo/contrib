package channel

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	Channel string      `md:"channel,required"` //The name of channel to use
	Data    interface{} `md:"data"`             //The data to put on the channel
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"channel": i.Channel,
		"data":    i.Data,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Channel, err = coerce.ToString(values["channel"])
	if err != nil {
		return err
	}
	i.Data = values["data"]

	return nil
}
