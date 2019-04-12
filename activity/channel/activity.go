package channel

import (
	"fmt"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/engine/channels"
)

func init() {
	_ = activity.Register(&Activity{})
}

var activityMd = activity.ToMetadata(&Input{})

// Activity is an Activity that is used to publish some data on a channel
// inputs : {channel, data}
// outputs: none
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval publishes the data  on the specified channel
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, err
	}

	if len(input.Channel) == 0 {
		return false, fmt.Errorf("channel name must be specified")
	}

	ch := channels.Get(input.Channel)

	if ch == nil {
		return false, fmt.Errorf("channel '%s' not registered with engine", input.Channel)
	}

	blocking := true

	//todo should we allow publish new wait?
	if blocking {
		ch.Publish(input.Data)
	} else {
		ch.PublishNoWait(input.Data)
	}

	if logger := ctx.Logger(); logger.DebugEnabled() {
		logger.Debugf("Published on '%s' value: %+v", input.Channel, input.Data)
	}

	return true, nil
}
