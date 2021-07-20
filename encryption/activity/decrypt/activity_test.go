package decrypt

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

	input := &Input{
		Data:       "1e236cb297e9af0ceae2ae50a90dac6eaacdb38a47c7a868a1cdfd1e29b8a9e9",
		Passphrase: "passphrase",
	}
	tc.SetInputObject(input)

	act.Eval(tc)

	output := &Output{}
	tc.GetOutputObject(output)
	assert.NotEmpty(t, output.Data)
}

func TestError(t *testing.T) {
	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	input := &Input{
		Data:       "aa1e236cb297e9af0ceae2ae50a90dac6eaacdb38a47c7a868a1cdfd1e29b8a9e9",
		Passphrase: "passphrase",
	}
	tc.SetInputObject(input)

	act.Eval(tc)

	output := &Output{}
	tc.GetOutputObject(output)
	assert.Empty(t, output.Data)
}
