package url

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFnQueryEncodeTrue(t *testing.T) {
	f := &fnQuery{}
	input := "https://subdomain.example.com/path?q=hello world#fragment with space"
	v, err := f.Eval(input, true)
	assert.Nil(t, err)

	assert.Equal(t, "q=hello+world", v)
}

func TestFnQueryEncodeFalse(t *testing.T) {
	f := &fnQuery{}
	input := "https://subdomain.example.com/path?q=hello world#fragment with space"
	v, err := f.Eval(input, false)
	assert.Nil(t, err)
	actualBytes, _ := json.Marshal(v)
	expectedMap := map[string][]string{
		"q": {"hello world"},
	}
	expectedBytes, _ := json.Marshal(expectedMap)
	assert.Equal(t, expectedBytes, actualBytes)
}
