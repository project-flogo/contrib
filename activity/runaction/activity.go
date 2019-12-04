package runaction

import (
	"context"
	"fmt"

	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/engine/runner"
	"github.com/project-flogo/core/support"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityMd = activity.ToMetadata(&Settings{}, &Output{})
var actionRunner = runner.NewDirect()

func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	ref := s.ActionRef

	if ref[0] == '#' {
		ref, _ = support.GetAliasRef("action", ref[1:])
	}

	factory := action.GetFactory(ref)
	if factory == nil {
		return nil, fmt.Errorf("unsupported action: %s", ref)
	}

	act, err := factory.New(&action.Config{Settings: s.ActionSettings})
	if err != nil {
		return nil, err
	}

	if act == nil {
		return nil, fmt.Errorf("unable to create action %s", ref)
	}

	md := act.IOMetadata()

	if md == nil && act.Metadata() != nil {
		md = act.Metadata().IOMetadata
	}

	var mdInput map[string]data.TypedValue

	if md != nil {
		mdInput = md.Input
	}

	return &Activity{settings: s, actionToRun: act, mdInput: mdInput}, nil
}

type Activity struct {
	settings    *Settings
	actionToRun action.Action
	mdInput     map[string]data.TypedValue
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	inputMap := make(map[string]interface{})

	for key, _ := range a.mdInput {
		inputMap[key] = ctx.GetInput(key)
	}

	result, err := actionRunner.RunAction(context.Background(), a.actionToRun, inputMap)
	if err != nil {
		return true, err
	}

	out := &Output{}
	out.Output = result

	err = ctx.SetOutputObject(out)
	if err != nil {
		return true, err
	}

	return true, nil
}
