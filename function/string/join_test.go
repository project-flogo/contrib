package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFnJoin_Eval(t *testing.T) {
	f := fnJoin{}
	var a = []string{"abc", "dddd", "cccc"}
	v, err := f.Eval(a, "-")
	assert.Nil(t, err)
	assert.Equal(t, "abc-dddd-cccc", v)
}
