package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFnPathEscape(t *testing.T) {
	f := &fnPathEscape{}
	input := "/some-path with ([brackets])"
	v, err := f.Eval(input)
	assert.Nil(t, err)
	assert.Equal(t, "%2Fsome-path%20with%20%28%5Bbrackets%5D%29", v)
}
