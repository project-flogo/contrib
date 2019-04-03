package appdata

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/app"
	"github.com/project-flogo/core/support/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestSet(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	settings := &Settings{Name: "test", Op: "set"}
	iCtx := test.NewActivityInitContext(settings, nil)

	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("value", "foo")

	_, err = act.Eval(tc)
	assert.Nil(t, err)

	appValue, _ := app.GetValue("test")
	assert.Equal(t, "foo", appValue)

}

func TestGet(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	err := app.SetValue("test", "bar")
	assert.Nil(t, err)

	settings := &Settings{Name: "test", Op: "get"}
	iCtx := test.NewActivityInitContext(settings, nil)

	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("value", "bar")

	_, err = act.Eval(tc)
	assert.Nil(t, err)

	appValue, _ := app.GetValue("test")
	assert.Equal(t, "bar", appValue)
}
