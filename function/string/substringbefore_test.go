package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var before = &Substringbefore{}

func TestStaticFunc_SubstringBefore(t *testing.T) {
	str := "TIBCO software Inc"
	final, _ := before.Eval(str, " ")
	fmt.Println(final)
	assert.Equal(t, final, "TIBCO")
}

func TestBeforeSample(t *testing.T) {
	final, _ := before.Eval("1999/04/01", "/")
	fmt.Println(final)
	assert.Equal(t, final, "1999")
}

func TestBeforeExpression(t *testing.T) {
	fun, err := factory.NewExpr(`string.substringBefore("1999/04/01","/")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}

func TestExpressionIPAS5021(t *testing.T) {
	fun, err := factory.NewExpr(`string.substringBefore("This is a sample server Petstore server ", "xxxx")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
