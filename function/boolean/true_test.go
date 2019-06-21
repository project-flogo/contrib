package boolean

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var s = &True{}

func TestStaticTrue(t *testing.T) {
	final1, _ := s.Eval(nil)
	fmt.Println(final1)
	assert.Equal(t, true, final1)
}

func TestTrueExpression(t *testing.T) {
	fun, err := factory.NewExpr(`boolean.true()`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)
}
