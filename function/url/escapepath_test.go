package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFnEscapePath(t *testing.T) {
	f := &fnEscapePath{}
	input := "https://subdomain.example.com/path?q=hello world#fragment with space"
	v, err := f.Eval(input)
	assert.Nil(t, err)
	assert.Equal(t, "/path", v)
}
