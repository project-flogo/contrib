package updobject

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	Object map[string]interface{} `md:"object"` // The object to update
	Values map[string]interface{} `md:"values"` // The key and values to update or add
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"object": i.Object,
		"values": i.Values,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error

	i.Object, err = coerce.ToObject(values["object"])
	if err != nil {
		return err
	}
	i.Values, err = coerce.ToObject(values["values"])

	return err
}

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityMd = activity.ToMetadata(&Input{})

// Activity is an Update Value Activity implementation
type Activity struct {
}

func New(ctx activity.InitContext) (activity.Activity, error) {
	return &Activity{}, nil
}

// Metadata implements activity.Activity.Metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements activity.Activity.Eval
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, err
	}


	for k, v := range input.Values {
		input.Object[k] = v
	}

	return true, nil
}
