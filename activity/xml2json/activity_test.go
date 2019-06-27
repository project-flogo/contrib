package xml2json

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

	aInput := &Input{XmlData: `<?xml version="1.0" encoding="UTF-8"?><hello>world</hello>`}
	tc.SetInputObject(aInput)
	done, _ := act.Eval(tc)
	assert.True(t,done)
	aOutput := &Output{}
    err := tc.GetOutputObject(aOutput)
    assert.Nil(t, err)
    assert.Equal(t, "world", aOutput.JsonObject["hello"])
}