package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var fn = &fnReplaceregex{}

func TestFnReplaceregex_Eval(t *testing.T) {
	v, err := fn.Eval("foo.*", "Tseafood", "People")
	assert.Nil(t, err)
	assert.Equal(t, "TseaPeople", v)
}
