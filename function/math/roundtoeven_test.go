package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFnRoundToEven(t *testing.T) {
	input := 12.50
	f := &fnRoundToEven{}
	v, err := f.Eval(input)
	assert.Nil(t, err)
	assert.Equal(t, 12.0, v)
}
