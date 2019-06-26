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
