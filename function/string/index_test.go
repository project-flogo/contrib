package string

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var index = &Index{}

func TestStatic_Index(t *testing.T) {
	final1, _ := index.Eval("TIBCO Web Integrator", "Web")
	fmt.Println(final1)
	assert.Equal(t, 6, final1)

	final2, _ := index.Eval("TIBCO Web Integrator", "Internet")
	fmt.Println(final2)
	assert.Equal(t, -1, final2)

	final3, _ := index.Eval("TIBCO 网络 Integrator", "Integrator")
	fmt.Println(final3)
	assert.Equal(t, 13, final3)
}

func TestIndexExpression(t *testing.T) {
	fun, err := factory.NewExpr(`string.index("TIBCO NAME", "NAME")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
