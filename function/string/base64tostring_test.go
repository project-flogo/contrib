package string

import (
	"fmt"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/data/expression/script"
	"github.com/project-flogo/core/data/resolve"
	"testing"

	"github.com/stretchr/testify/assert"
)

var base64S = &Base64ToString{}

var resolver = resolve.NewCompositeResolver(map[string]resolve.Resolver{"static": nil, ".": nil, "env": &resolve.EnvResolver{}})
var factory = script.NewExprFactory(resolver)

func init() {
	function.ResolveAliases()
}

func TestStaticFunc_Base64_to_string(t *testing.T) {
	final1, err := base64S.Eval("SGVsbG8sIOS4lueVjA==")
	fmt.Println(final1)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, "Hello, 世界", final1)

	//Negative test case: invalid input base64 string
	final2, err := base64S.Eval("SSSSGVsbG8sIOS4lueVjA==")
	fmt.Println(final2)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, "", final2)

}

func TestBase64Expression(t *testing.T) {
	fun, err := factory.NewExpr(`string.base64ToString("SGVsbG8sIOS4lueVjA==")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, "Hello, 世界", v)
}
