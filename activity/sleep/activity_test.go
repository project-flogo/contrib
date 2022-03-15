package sleep

import (
	"github.com/project-flogo/core/support/test"
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	input := &Input{SleepTime: 1}
	tc.SetInputObject(input)

	act.Eval(tc)
}

func TestAddToFlow(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	//setup attrs
	tc.SetInput("sleepTime", 1)
	act.Eval(tc)
}
