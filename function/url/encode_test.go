package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFnEncode(t *testing.T) {
	f := &fnEncode{}
	input := "https://subdomain.example.com/path to be escaped?q=hello world#fragment with space"
	v, err := f.Eval(input)
	assert.Nil(t, err)
	expected := "https://subdomain.example.com/path%20to%20be%20escaped?q=hello+world#fragment%20with%20space"
	assert.Equal(t, expected, v)
}
