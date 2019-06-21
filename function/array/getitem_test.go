package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var g = &Get{}

func TestStaticGet(t *testing.T) {
	expectedResult := []string{"Cat", "Dog", "Snake"}
	final, err := g.Eval(expectedResult, 1)
	assert.Nil(t, err)
	assert.Equal(t, "Dog", final)

	expectedResult = []string{"Cat", "Dog", "Snake"}
	_, err = g.Eval(expectedResult, 4)
	assert.NotNil(t, err)
}

func TestGetExpression(t *testing.T) {
	fun, err := factory.NewExpr(`array.get(array.create("123","456"), 1)`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, "456", v)
}
