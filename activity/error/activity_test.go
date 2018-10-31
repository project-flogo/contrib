package error

import (
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestSimpleError(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	input := &Input{Message: "test error"}
	tc.SetInputObject(input)

	//eval
	done, err := act.Eval(tc)
	assert.False(t, done)
	assert.NotNil(t, err)

	ae, ok := err.(*activity.Error)
	assert.True(t, ok)
	assert.Equal(t, "test error", ae.Error())
}
