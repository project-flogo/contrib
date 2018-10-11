package string

import (
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFnSubstring_Eval(t *testing.T) {
	f := &fnSubstring{}
	v, err := function.Eval(f, "abc", 1, -1)
	assert.Nil(t, err)
	assert.Equal(t, "bc", v)

	v, err = function.Eval(f, "abc", 1, 1)
	assert.Nil(t, err)
	assert.Equal(t, "b", v)
}
