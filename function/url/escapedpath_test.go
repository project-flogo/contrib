package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFnEscapedPath(t *testing.T) {
	f := &fnEscapedPath{}
	input := "https://example.com:8080/root-path/sub%2Fpath?query=example+query+%2F+question#fragment"
	v, err := f.Eval(input)
	assert.Nil(t, err)
	assert.Equal(t, "/root-path/sub%2Fpath", v)
}
