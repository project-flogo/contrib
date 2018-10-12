package activity_return

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/logger"
)

// log is the default package logger
var log = logger.GetLogger("activity-return")

func init() {
	activity.Register(&ReturnActivity{}, New)
}

type Settings struct {
	Mappings map[string]interface{} `md:"mappings"`
}

var activityMd = activity.ToMetadata(&Settings{})

func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	act := &ReturnActivity{}

	log.Debugf("Mappings: %+v", s.Mappings)

	act.mapper, err = ctx.MapperFactory().NewMapper(s.Mappings)
	if err != nil {
		return nil, err
	}

	return act, nil
}

// ReturnActivity is an Activity that is used to return/return via the trigger
// inputs : {method,uri,params}
// outputs: {result}
type ReturnActivity struct {
	mapper mapper.Mapper
}

func (a *ReturnActivity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Invokes a REST Operation
func (a *ReturnActivity) Eval(ctx activity.Context) (done bool, err error) {

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
