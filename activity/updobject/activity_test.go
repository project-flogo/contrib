package updobject

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestAdd(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	iCtx := test.NewActivityInitContext(nil, nil)

	act, err := New(iCtx)
	assert.Nil(t, err)

	obj := make(map[string]interface{})

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("object", obj)
	tc.SetInput("values", map[string]interface{}{"foo":"bar"})

	_, err = act.Eval(tc)
	assert.Nil(t, err)

	v, exists := obj["foo"]
	assert.True(t, exists)
	assert.Equal(t, "bar", v)
}


func TestSet(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	iCtx := test.NewActivityInitContext(nil, nil)

	act, err := New(iCtx)
	assert.Nil(t, err)

	obj := make(map[string]interface{})
	obj["foo"] = "bar"

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("object", obj)
	tc.SetInput("values", map[string]interface{}{"foo":"bar2"})

	_, err = act.Eval(tc)
	assert.Nil(t, err)

	assert.Equal(t, "bar2", obj["foo"])
}

