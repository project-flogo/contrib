package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFnCeil(t *testing.T) {
	input := 1.49
	f := &fnCeil{}
	v, err := f.Eval(input)
	assert.Nil(t, err)
	assert.Equal(t, 2.0, v)
}
