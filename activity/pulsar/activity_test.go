package sample

/*
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
func TestSettings(t *testing.T) {
	settings := &Settings{URL: "pulsar://localhost:6650", Topic: "sample"}

	iCtx := test.NewActivityInitContext(settings, nil)
	act, err := New(iCtx)
	assert.Nil(t, err)
	assert.NotNil(t, act)

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("payload", "sample")

	_, err = act.Eval(tc)
	assert.Nil(t, err)
}
*/
