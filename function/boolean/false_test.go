package boolean

import (
	"fmt"
	"github.com/project-flogo/core/data/expression/script"
	"github.com/project-flogo/core/data/resolve"
	"github.com/stretchr/testify/assert"
	"testing"
)

var resolver = resolve.NewCompositeResolver(map[string]resolve.Resolver{"static": nil, ".": nil, "env": &resolve.EnvResolver{}})
var factory = script.NewExprFactory(resolver)

var f = &False{}

func TestStaticFalse(t *testing.T) {
	final1, _ := f.Eval(nil)
	fmt.Println(final1)
	assert.Equal(t, false, final1)
}

func TesFalsetExpression(t *testing.T) {
	fun, err := factory.NewExpr(`boolean.false()`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v)
}
