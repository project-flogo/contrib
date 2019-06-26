package string

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

var eq = &fnEquals{}

func TestFnEquals_Eval(t *testing.T) {

	v, err := function.Eval(eq, "foo", "bar")
	assert.Nil(t, err)
	assert.False(t, v.(bool))

	v, err = function.Eval(eq, "foo", "foo")
	assert.Nil(t, err)
	assert.True(t, v.(bool))
}

func TestStaticFunc_Eq(t *testing.T) {
	final1, _ := eq.Eval("TIBCO Web", "TIBCO")
	fmt.Println(final1)
	assert.Equal(t, false, final1)

	final2, _ := eq.Eval("123T", "123t")
	fmt.Println(final2)
	assert.Equal(t, false, final2)

}
