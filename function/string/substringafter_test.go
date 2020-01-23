package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sub = &Substringafter{}

func TestStaticFunc_SubstringAfter(t *testing.T) {
	str := "TIBCO software Inc"
	final, _ := sub.Eval(str, " ")
	assert.Equal(t, final, "software Inc")
}

func TestSubStringAfterSample(t *testing.T) {
	final, _ := sub.Eval("1999/04/01", "/")
	assert.Equal(t, final, "04/01")
}
