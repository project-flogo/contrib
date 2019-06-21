package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sp = &Split{}

func TestStaticFunc_Split(t *testing.T) {
	final1, _ := sp.Eval("TIBCO Web Integrator", " ")
	final := final1.([]string)
	assert.Equal(t, "TIBCO", final[0])
	assert.Equal(t, "Web", final[1])
	assert.Equal(t, "Integrator", final[2])

	final2, _ := sp.Eval("TIBCO。网路。Integrator", "。")
	fmt.Println(final2)
	final = final2.([]string)

	assert.Equal(t, "TIBCO", final[0])
	assert.Equal(t, "网路", final[1])
	assert.Equal(t, "Integrator", final[2])

}

func TestSplitExpression(t *testing.T) {
	fun, err := factory.NewExpr(`string.split("seafood,name", ",")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}

func TestSplitExpression2(t *testing.T) {
	fun, err := factory.NewExpr(`string.split("seafood namefood", " ")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
