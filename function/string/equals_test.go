package string

import (
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFnEquals_Eval(t *testing.T) {
	f := &fnEquals{}

	v, err := function.Eval(f, "foo", "bar")
	assert.Nil(t, err)
	assert.False(t, v.(bool))

	v, err = function.Eval(f, "foo", "foo")
	assert.Nil(t, err)
	assert.True(t, v.(bool))
}
