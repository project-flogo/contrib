package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFnTitle_Eval(t *testing.T) {
	f := fnTitle{}
	str := "hello world"
	v, err := f.Eval(str)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World", v)
}
