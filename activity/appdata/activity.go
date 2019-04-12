package appdata

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/app"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
)

const (
	opGet int = 0
	opSet int = 1

	ivValue = "value"
	ovValue = "value"
)

type Settings struct {
	Name string `md:"name,required"`       // The name of the shared attribute
	Op   string `md:"op,allowed(get,set)"` // The operation (get or set), 'get' is the default
	Type string `md:"type"`                // The data type of the shared value, default is 'any'
}

type Input struct {
	Value interface{} `md:"value"` // The value of the shared attribute
}

type Output struct {
	Value interface{} `md:"value"` // The value of the shared attribute
}

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityMd = activity.ToMetadata(&Settings{}, &Output{})

// Activity is a Counter Activity implementation
type Activity struct {
	op       int
	dt       data.Type
	attrName string
}

func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	act := &Activity{attrName: s.Name}

	if s.Op == "set" {
		act.op = opSet
	}

	if s.Type != "" {
		t, err := data.ToTypeEnum(s.Type)
		if err != nil {
			return nil, err
		}
		act.dt = t
	}

	return act, nil
}

// Metadata implements activity.Activity.Metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements activity.Activity.Eval
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	switch a.op {
	case opGet:

		val, exists := app.GetValue(a.attrName)
		if exists && a.dt > 1 {
			val, err = coerce.ToType(val, a.dt)
			if err != nil {
				return false, err
			}
		}
		err = ctx.SetOutput(ovValue, val)
		if err != nil {
			return false, err
		}
	case opSet:
		val := ctx.GetInput(ivValue)

		if a.dt > 1 {
			val, err = coerce.ToType(val, a.dt)
			if err != nil {
				return false, err
			}
		}

		err = app.SetValue(a.attrName, val)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}
