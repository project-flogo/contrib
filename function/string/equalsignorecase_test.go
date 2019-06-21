package string

import (
	"fmt"
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
	fmt.Println(final1)
	assert.Equal(t, false, final1)

	final2, _ := eqi.Eval("TIBCO", "tibco")
	fmt.Println(final2)
	assert.Equal(t, true, final2)

}

func TestEQIExpression(t *testing.T) {
	fun, err := factory.NewExpr(`string.equalsIgnoreCase("TIBCO FLOGO", "TIBCO FLOGO")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	assert.Equal(t, true, v)
}

func TestExpressionIgnoreCase(t *testing.T) {
	fun, err := factory.NewExpr(`string.equalsIgnoreCase("TIBCO flogo", "TIBCO FLOGO")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	assert.Equal(t, true, v)
}
