package validatehash

import "github.com/project-flogo/core/data/coerce"

// Settings for the activity
type Settings struct {
}

// Input for the activity
type Input struct {
	Hash string `md:"hash"` // Data for the hash
	Data string `md:"data"` // LiveApps element: allowed values are action and workitem
}

// ToMap for Input
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"hash": i.Hash,
		"data": i.Data,
	}
}

// FromMap for Input
func (i *Input) FromMap(values map[string]interface{}) error {
	i.Hash, _ = coerce.ToString(values["hash"])
	i.Data, _ = coerce.ToString(values["data"])

	return nil
}

// Output for the activity
type Output struct {
	Valid bool `md:"valid"`
}

// ToMap conver to object
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"valid": o.Valid,
	}
}

// FromMap convert to object
func (o *Output) FromMap(values map[string]interface{}) error {
	o.Valid, _ = coerce.ToBool(values["valid"])
	return nil
}
