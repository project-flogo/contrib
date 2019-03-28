package error

import (
	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{})
}

var activityMd = activity.ToMetadata(&Input{})

// Activity is an Activity that used to cause an explicit error in the flow
// inputs : {message,data}
// outputs: node
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval returns an error
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, err
	}

	if logger := ctx.Logger(); logger.DebugEnabled() {
		logger.Debugf("Message :'%s', Data: '%+v'", input.Message, input.Data)
	}

	return false, activity.NewError(input.Message, "", input.Data)
}
