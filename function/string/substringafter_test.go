package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sub = &Substringafter{}

func TestStaticFunc_SubstringAfter(t *testing.T) {
	str := "TIBCO software Inc"
	final, _ := sub.Eval(str, " ")
	fmt.Println(final)
	assert.Equal(t, final, "software Inc")
}

func TestSubStringAfterSample(t *testing.T) {
	final, _ := sub.Eval("1999/04/01", "/")
	fmt.Println(final)
	assert.Equal(t, final, "04/01")
}

func TestSubStringAfterExpression(t *testing.T) {
	fun, err := factory.NewExpr(`string.substringAfter("1999/04/01","/")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}

func TestExpressionIPAS4962(t *testing.T) {
	fun, err := factory.NewExpr(`string.substringAfter("This is a sample server Petstore server ","sample")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}

func TestSubStringExpressionIPAS5021(t *testing.T) {
	fun, err := factory.NewExpr(`string.substringAfter("This is a sample server Petstore server ", "This")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
