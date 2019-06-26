package array

import (
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
