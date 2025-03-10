package sleep

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"

	"time"
)

type Input struct {
	SleepTime int `md:"SleepTime,required"`
}

type Activity struct {
}

func init() {
	_ = activity.Register(&Activity{})
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"sleepTime": i.SleepTime,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error
	i.SleepTime, err = coerce.ToInt(values["sleepTime"])
	if err != nil {
		return err
	}

	return nil
}

func (a *Activity) Metadata() *activity.Metadata {
	return activity.ToMetadata(&Input{})
}

func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	ctx.GetInputObject(input)

	interval := input.SleepTime
	time.Sleep(time.Duration(interval) * time.Second)

	return true, nil
}
