package array

import (
	"fmt"
	"github.com/project-flogo/core/data/expression/function"
	"testing"

	"github.com/stretchr/testify/assert"
)

var c = &Count{}

func TestStaticCount(t *testing.T) {
	expectedResult := []string{"Cat", "Dog", "Snake"}
	final, err := c.Eval(expectedResult)
	assert.Nil(t, err)
	fmt.Println(final)
	assert.Equal(t, 3, final)
}

func TestCountExpression(t *testing.T) {
	function.ResolveAliases()
	fun, err := factory.NewExpr(`array.count(array.create("123","456"))`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, 2, v)
}
