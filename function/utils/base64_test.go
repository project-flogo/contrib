package utils

import (
	"github.com/project-flogo/core/data/coerce"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFnDecodeBase64_Eval(t *testing.T) {
	f := fnDecodeBase64{}
	v, err := f.Eval("SGVsbG8sIFdvcmxk")
	assert.Nil(t, err)
	vv, _ := coerce.ToString(v)
	assert.Equal(t, "Hello, World", vv)
}

func TestFnEncodeBase64_Eval(t *testing.T) {
	f := fnEncodeBase64{}
	v, err := f.Eval([]byte("Hello, World"))
	assert.Nil(t, err)
	vv, _ := coerce.ToString(v)
	assert.Equal(t, "SGVsbG8sIFdvcmxk", vv)
}
