package array

import (
	"fmt"
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
