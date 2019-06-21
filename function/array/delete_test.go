package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var d = &Delete{}

func TestStaticDelete(t *testing.T) {
	expectedResult := []string{"Cat", "Dog", "Snake"}
	final, err := d.Eval(expectedResult, 2)
	assert.Nil(t, err)
	fmt.Println(final)
	assert.Equal(t, []string{"Cat", "Dog"}, final)
}

func TestDelateExpression(t *testing.T) {
	fun, err := factory.NewExpr(`array.delete(array.create("123","456"), 1)`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, []interface{}{"123"}, v)
}
