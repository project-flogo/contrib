package obtainhash

import "github.com/project-flogo/core/data/coerce"

// Settings for the activity
type Settings struct {
}

// Input for the activity
type Input struct {
	Data string `md:"data"` // Data for the hash
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"data": i.Data,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {
	i.Data, _ = coerce.ToString(values["data"])

	return nil
}

// Output for the activity
type Output struct {
	Hash string `md:"hash"`
}

// ToMap conver to object
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"hash": o.Hash,
	}
}

// FromMap convert to object
func (o *Output) FromMap(values map[string]interface{}) error {

	o.Hash, _ = coerce.ToString(values["hash"])

	return nil
}
