package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var end = &EndsWith{}

func TestStaticEndWith(t *testing.T) {
	final1, _ := end.Eval("TIBCO Web Integrator", "Integrator")
	fmt.Println(final1)
	assert.Equal(t, true, final1)

	final2, _ := end.Eval("TIBCO Web 集成器", "集成器")
	fmt.Println(final2)
	assert.Equal(t, true, final2)

	final3, _ := end.Eval("TIBCO 网路 Integrator", "网路")
	fmt.Println(final3)
	assert.Equal(t, false, final3)
}

func TestEndWithExpression(t *testing.T) {
	fun, err := factory.NewExpr(`string.endsWith("TIBCO NAME", "NAME")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
