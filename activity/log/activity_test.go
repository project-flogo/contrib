package log

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

func TestEval(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	input := &Input{Message: "test message", AddDetails: true}
	tc.SetInputObject(input)

	act.Eval(tc)
}

func TestAddToFlow(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	//setup attrs
	tc.SetInput("message", "test message")
	tc.SetInput("addDetails", true)

	act.Eval(tc)
}
