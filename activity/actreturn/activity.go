package actreturn

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

type Settings struct {
	Mappings map[string]interface{} `md:"mappings"`  // Set of mappings to execute when the activity runs
}

var activityMd = activity.ToMetadata(&Settings{})

func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	act := &Activity{}

	ctx.Logger().Debugf("Mappings: %+v", s.Mappings)

	act.mapper, err = ctx.MapperFactory().NewMapper(s.Mappings)
	if err != nil {
		return nil, err
	}

	return act, nil
}

// Activity is an Activity that is used to return/return via the trigger
// inputs : {method,uri,params}
// outputs: {result}
type Activity struct {
	mapper mapper.Mapper
}

func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Invokes a REST Operation
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	actionCtx := ctx.ActivityHost()

	if a.mapper == nil {
		//No mapping
		actionCtx.Return(nil, nil)
		return true, nil
	}

	inputScope := actionCtx.Scope() //host data

	results, err := a.mapper.Apply(inputScope)
	if err != nil {
		return false, activity.NewError(err.Error(), "", nil)
	}

	actionCtx.Return(results, nil)

	return true, nil
}
