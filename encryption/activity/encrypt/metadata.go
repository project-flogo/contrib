package encrypt

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
}

// Input for the activity
type Input struct {
	Data       string `md:"data"`       // Data to encrypt
	Passphrase string `md:"passphrase"` // Encryption passphrase
}

// ToMap for Input
func (i *Input) ToMap() map[string]interface{} {

	return map[string]interface{}{
		"data":       i.Data,
		"passphrase": i.Passphrase,
	}
}

// FromMap for input
func (i *Input) FromMap(values map[string]interface{}) error {
	i.Data, _ = coerce.ToString(values["data"])
	i.Passphrase, _ = coerce.ToString(values["passphrase"])

	return nil
}

// Output for the activity
type Output struct {
	Data string `md:"data"` // Data encrypted
}

// ToMap conver to object
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"data": o.Data,
	}
}

// FromMap convert to object
func (o *Output) FromMap(values map[string]interface{}) error {
	o.Data, _ = coerce.ToString(values["data"])

	return nil
}
