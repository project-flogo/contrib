package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFnEscapeQuery(t *testing.T) {
	f := &fnQueryEscape{}
	input := "hello world"
	v, err := f.Eval(input)
	assert.Nil(t, err)
	assert.Equal(t, "hello+world", v)
}
