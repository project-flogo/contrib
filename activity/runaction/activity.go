package runaction

import (
	"context"
	"errors"
	"fmt"

	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support"
	"github.com/project-flogo/core/runner"
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

	act := &Activity{settings: s}

	//ctx.Logger().Debugf("flowURI: %+v", s.FlowURI)

	return act, nil
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

// Activity is an Activity that is used to log a message to the console
// inputs : {message, flowInfo}
// outputs: none
type Activity struct {
	settings *Settings
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	out := &Output{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	ref, _ := support.GetAliasRef("action", a.settings.Ref[1:])

	factory := action.GetFactory(ref)

	var act action.Action
	settingsURI := make(map[string]interface{})

	settingsURI[settings.UriType] = a.settings.ResURI //a.settings.ResURI

	act, err = factory.New(&action.Config{Settings: settingsURI})

	if err != nil || act == nil {
		ctx.Logger().Infof("Error in Inialtization of Sync Action %v", err)
		return false, err
	}
	inputMap := make(map[string]interface{})
	
	_, isMap := input.Input.(map[string]interface{})
	if !isMap {
		inputMap["input"] = input.Input
	}
	
	engineRunner := runner.NewDirect()
	
	var result map[string]interface{}

	if !isMap {
		result, err = engineRunner.RunAction((context.Background(), act ,inputMap)
	} else {
		result, err = engineRunner.RunAction((context.Background(), act ,input.Input.(map[string]interface{}))
	}

	if err != nil {
		ctx.Logger().Infof("Error in Running  Action %v", err)
		return true, fmt.Errorf("Error in Running Action: %v", err)
	}

	out.Output = result

	ctx.SetOutputObject(out)

	return true, nil

	
}
