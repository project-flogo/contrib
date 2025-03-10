package validatehash

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

	input := &Input{Hash: "3f0a377ba0a4a460ecb616f6507ce0d8cfa3e704025d4fda3ed0c5ca05468728", Data: "test message"}
	tc.SetInputObject(input)

	act.Eval(tc)

	output := &Output{}
	tc.GetOutputObject(output)

	assert.True(t, output.Valid)
}
