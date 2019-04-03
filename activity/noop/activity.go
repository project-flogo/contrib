package noop

import (
	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{})
}

var activityMd = activity.ToMetadata()

// Activity is an Activity that is used to log a message to the console
// inputs : none
// outputs: none
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	if ctx.Logger().DebugEnabled() {
		ctx.Logger().Debug("Performing No-Op Activity")
	}

	return true, nil
}
