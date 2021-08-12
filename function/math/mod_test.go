package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFnMod(t *testing.T) {
	input1 := 7.0
	input2 := 4.0
	f := &fnMod{}
	v, err := f.Eval(input1, input2)
	assert.Nil(t, err)
	assert.Equal(t, 3.0, v)
}
