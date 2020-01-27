package string

import (
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFnEqualsIgnoreCase_Eval(t *testing.T) {
	f := &fnEqualsIgnoreCase{}

	v, err := function.Eval(f, "foo", "Bar")
	assert.Nil(t, err)
	assert.False(t, v.(bool))

	v, err = function.Eval(f, "foo", "Foo")
	assert.Nil(t, err)
	assert.True(t, v.(bool))
}

var eqi = &fnEqualsIgnoreCase{}

func TestStaticFuncEQI(t *testing.T) {
	final1, _ := eqi.Eval("TIBCO FLOGO", "TIBCO")

	assert.Equal(t, false, final1)

	final2, _ := eqi.Eval("TIBCO", "tibco")

	assert.Equal(t, true, final2)

}
