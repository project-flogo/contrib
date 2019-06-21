package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sTbase = &StringToBase64{}

func TestStaticFunc_String_to_base64(t *testing.T) {
	final1, _ := sTbase.Eval("Hello, 世界")
	fmt.Println(final1)
	assert.Equal(t, "SGVsbG8sIOS4lueVjA==", final1)

}

func TestStringBase64Expression(t *testing.T) {
	fun, err := factory.NewExpr(`string.stringToBase64("Hello, 世界")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
