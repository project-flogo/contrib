package activity_reply

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/logger"
)

// log is the default package logger
var log = logger.GetLogger("activity-reply")

func init() {
	activity.Register(&ReplyActivity{}, New)
}

type Settings struct {
	Mappings map[string]interface{} `md:"mappings,required"`
}

var activityMd = activity.ToMetadata(&Settings{})

func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	act := &ReplyActivity{}

	log.Debugf("Mappings: %+v", s.Mappings)

	act.mapper, err = ctx.MapperFactory().NewMapper(s.Mappings)
	if err != nil {
		return nil, err
	}

	return act, nil
}

// ReplyActivity is an Activity that is used to reply/return via the trigger
// inputs : {method,uri,params}
// outputs: {result}
type ReplyActivity struct {
	mapper mapper.Mapper
}

func (a *ReplyActivity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Invokes a REST Operation
func (a *ReplyActivity) Eval(ctx activity.Context) (done bool, err error) {

	actionCtx := ctx.ActivityHost()

	if a.mapper == nil {
		//No mapping
		actionCtx.Reply(nil, nil)
		return true, nil
	}

	inputScope := actionCtx.Scope() //host data

	results, err := a.mapper.Apply(inputScope)
	if err != nil {
		return false, activity.NewError(err.Error(), "", nil)
	}

	actionCtx.Reply(results, nil)

	return true, nil
}
