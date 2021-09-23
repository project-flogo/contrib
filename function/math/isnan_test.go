package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFnIsNaN(t *testing.T) {
	input := 1.0
	f := &fnIsNaN{}
	v, err := f.Eval(input)
	assert.Nil(t, err)
	assert.Equal(t, false, v)
}
