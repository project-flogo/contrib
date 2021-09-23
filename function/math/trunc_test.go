package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFnTrunc(t *testing.T) {
	input := 3.142
	f := &fnTrunc{}
	v, err := f.Eval(input)
	assert.Nil(t, err)
	assert.Equal(t, 3.0, v)
}
