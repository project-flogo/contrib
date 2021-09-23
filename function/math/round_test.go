package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFnRoundDown(t *testing.T) {
	input := 1.49
	f := &fnRound{}
	v, err := f.Eval(input)
	assert.Nil(t, err)
	assert.Equal(t, 1.0, v)
}

func TestFnRoundUp(t *testing.T) {
	input := 1.50
	f := &fnRound{}
	v, err := f.Eval(input)
	assert.Nil(t, err)
	assert.Equal(t, 2.0, v)
}
