package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFnFloor(t *testing.T) {
	input := 1.51
	f := &fnFloor{}
	v, err := f.Eval(input)
	assert.Nil(t, err)
	assert.Equal(t, 1.0, v)
}
