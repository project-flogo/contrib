package channel

import (
	"testing"
	"time"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/engine/channels"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	channels.New("test", 5)
	ch := channels.Get("test")

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	input := &Input{Channel: "test", Data: 2}
	tc.SetInputObject(input)

	done, err := act.Eval(tc)
	assert.Nil(t, err)
	assert.True(t, done)

	var found interface{}

	ch.RegisterCallback(func(msg interface{}) {
		found = msg
	})

	channels.Start()

	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, 2, found)

	channels.Stop()
}
