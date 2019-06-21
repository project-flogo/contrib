package array

import (
	"fmt"
	"github.com/project-flogo/core/data/expression/function"
	"testing"

	"github.com/stretchr/testify/assert"
)

var create = &Create{}

func TestStaticFunc_ArrayString(t *testing.T) {
	expectedResult := []string{"Cat", "Dog", "Snake"}
	final, err := create.Eval("Cat", "Dog", "Snake")
	assert.Nil(t, err)
	for i, item := range final.([]interface{}) {
		assert.Equal(t, item.(string), expectedResult[i])
	}
}

func TestExpression(t *testing.T) {
	function.ResolveAliases()
	expectedResult := []string{"123", "456"}
	fun, err := factory.NewExpr(`array.create("123","456")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	for i, item := range v.([]interface{}) {
		assert.Equal(t, item.(string), expectedResult[i])
	}
}

func TestExpressionDifferentType(t *testing.T) {
	function.ResolveAliases()
	fun, err := factory.NewExpr(`array.create(123,456.33)`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
	assert.Equal(t, 123, v.([]interface{})[0])
	assert.Equal(t, 456.33, v.([]interface{})[1])

	fun, err = factory.NewExpr(`array.create("123",456.33)`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err = fun.Eval(nil)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
	assert.Equal(t, "123", v.([]interface{})[0])
	assert.Equal(t, 456.33, v.([]interface{})[1])

}

func TestExpression2(t *testing.T) {
	function.ResolveAliases()

	fun, err := factory.NewExpr(`array.create("adi","shukla",false)`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)

	assert.Equal(t, `shukla`, v.([]interface{})[1])

}

func TestExpression3(t *testing.T) {
	function.ResolveAliases()

	fun, err := factory.NewExpr(`array.create("adi","shukla",true)`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)

	assert.Equal(t, `shukla`, v.([]interface{})[1])
}
