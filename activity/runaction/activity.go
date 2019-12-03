package runaction

import (
	"context"

	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/engine/runner"
	"github.com/project-flogo/core/support"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	ref, _ := support.GetAliasRef("action", s.ActionRef[1:])

	factory := action.GetFactory(ref)

	act, err := factory.New(&action.Config{Settings: s.ActionSettings})

	if err != nil || act == nil {
		ctx.Logger().Infof("Error in Inialtization of Sync Action %v", err)
		return nil, err
	}

	return &Activity{settings: s, action: act}, nil
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

type Activity struct {
	settings *Settings
	action   action.Action
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	out := &Output{}

	inputMap := make(map[string]interface{})

	for key, _ := range a.action.IOMetadata().Input {
		inputMap[key] = ctx.GetInput(key)
	}

	engineRunner := runner.NewDirect()

	result, err := engineRunner.RunAction(context.Background(), a.action, inputMap)

	if err != nil {
		ctx.Logger().Infof("Error in Running  Action %v", err)
		return true, err
	}

	out.Output = result

	ctx.SetOutputObject(out)

	return true, nil

}
