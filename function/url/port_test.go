package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFnPort(t *testing.T) {
	f := &fnPort{}
	input := "https://subdomain.example.com:8080/path?q=hello world#fragment with space"
	v, err := f.Eval(input)
	assert.Nil(t, err)
	assert.Equal(t, "8080", v)
}
