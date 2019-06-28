package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var contain = &Contains{}

func TestStaticFunc_Contains(t *testing.T) {

	array := []string{"Cat", "Dog", "Snake"}
	final, _ := contain.Eval(array, "Snake")
	assert.Equal(t, true, final)
	final, _ = contain.Eval(array, "Foo")
	assert.Equal(t, false, final)

	arrayInt := []int{5, 40, 10}
	final, _ = contain.Eval(arrayInt, 40)
	assert.Equal(t, true, final)
	final, _ = contain.Eval(arrayInt, "Foo")
	assert.Equal(t, false, final)

}
