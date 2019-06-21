package string

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFnContains_Eval(t *testing.T) {
	f := &fnContains{}

	v, err := function.Eval(f, "foo", "Bar")
	assert.Nil(t, err)
	assert.False(t, v.(bool))

	v, err = function.Eval(f, "foobar", "foo")
	assert.Nil(t, err)
	assert.True(t, v.(bool))
}

func TestContainsExpression(t *testing.T) {
	fun, err := factory.NewExpr(`string.contains("TIBCO Web Integrator","Web")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)
	fmt.Println(v)
}
