package number

import (
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/data/expression/script"
	"github.com/project-flogo/core/data/resolve"
	"github.com/stretchr/testify/assert"
	"testing"
)

var s = &fnRandom{}

var resolver = resolve.NewCompositeResolver(map[string]resolve.Resolver{"static": nil, ".": nil, "env": &resolve.EnvResolver{}})
var factory = script.NewExprFactory(resolver)

func init() {
	function.ResolveAliases()
}

func TestSample(t *testing.T) {
	final1, _ := s.Eval(100)
	assert.NotNil(t, final1)
}

func TestRandomExpression(t *testing.T) {
	fun, err := factory.NewExpr(`number.random(10)`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.NotNil(t, v)
}
