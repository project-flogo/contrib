package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFnHostname(t *testing.T) {
	f := &fnHostname{}
	input := "https://subdomain.example.com/path?q=hello world#fragment with space"
	v, err := f.Eval(input)
	assert.Nil(t, err)
	assert.Equal(t, "subdomain.example.com", v)
}
