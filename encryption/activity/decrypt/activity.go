package decrypt

import (
	"github.com/project-flogo/contrib/encryption/crypt"
	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

// New function for the activity
func New(ctx activity.InitContext) (activity.Activity, error) {
	act := &Activity{}

	return act, nil
}

// Activity is an activity that is used to invoke a REST Operation
type Activity struct {
}

// Metadata for the activity
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Create the hash
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		err = activity.NewError(err.Error(), "DECRYPT-001", nil)
		return false, err
	}

	logger := ctx.Logger()

	if logger.DebugEnabled() {
		logger.Debugf("Input params: %s", input)
	}
	data, err := crypt.DecryptString(input.Data, input.Passphrase)
	if err != nil {
		err = activity.NewError(err.Error(), "DECRYPT-002", nil)
		return false, err
	}

	output := &Output{Data: data}
	err = ctx.SetOutputObject(output)
	if err != nil {
		err = activity.NewError(err.Error(), "DECRYPT-003", nil)
		return false, err
	}

	return true, nil
}
