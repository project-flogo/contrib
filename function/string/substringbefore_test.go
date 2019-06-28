package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var before = &Substringbefore{}

func TestStaticFunc_SubstringBefore(t *testing.T) {
	str := "TIBCO software Inc"
	final, _ := before.Eval(str, " ")
	fmt.Println(final)
	assert.Equal(t, final, "TIBCO")
}

func TestBeforeSample(t *testing.T) {
	final, _ := before.Eval("1999/04/01", "/")
	fmt.Println(final)
	assert.Equal(t, final, "1999")
}
