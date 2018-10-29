package counter

import (
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/resolve"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestIncrement(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	settings := &Settings{CounterName: "test", Op: "increment"}
	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(settings, mf)

	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	act.Eval(tc)

	value := tc.GetOutput(ovValue).(int)

	assert.Equal(t, 1, value)
}

func TestGet(t *testing.T) {

	settings := &Settings{CounterName: "test", Op: "get"}
	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(settings, mf)

	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	c := counters["test"]
	c.Reset()
	c.Increment()
	c.Increment()
	c.Increment()

	act.Eval(tc)

	value := tc.GetOutput(ovValue).(int)

	assert.Equal(t, 3, value)
}

func TestReset(t *testing.T) {

	settings := &Settings{CounterName: "test", Op: "reset"}
	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(settings, mf)

	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	c := counters["test"]
	c.Reset()
	c.Increment()
	c.Increment()
	c.Increment()

	act.Eval(tc)

	value := tc.GetOutput(ovValue).(int)

	assert.Equal(t, 0, value)
}
