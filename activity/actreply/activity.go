package actreply

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/metadata"
)

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

	ctx.Logger().Debugf("Mappings: %+v", s.Mappings)

	act.mapper, err = ctx.MapperFactory().NewMapper(s.Mappings)
	if err != nil {
		return nil, err
	}

	return act, nil
}

// ReplyActivity is an ReplyActivity that is used to reply/return via the trigger
// inputs : {method,uri,params}
// outputs: {result}
type ReplyActivity struct {
	mapper mapper.Mapper
}

func (a *ReplyActivity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.ReplyActivity.Eval - Invokes a REST Operation
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
