package error

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	Message string      `md:"message"` // The error message
	Data    interface{} `md:"data"`    // The error data
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"message": i.Message,
		"data":    i.Data,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Message, err = coerce.ToString(values["message"])
	if err != nil {
		return err
	}
	i.Data = values["data"]

	return nil
}
