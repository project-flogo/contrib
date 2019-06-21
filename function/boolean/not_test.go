package boolean

import (
	"fmt"
	"github.com/project-flogo/core/data/expression/function"
	"os"

	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var bb = &Not{}

func TestMain(m *testing.M) {
	function.ResolveAliases()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestStaticFNots(t *testing.T) {
	final1, _ := bb.Eval(true)
	fmt.Println(final1)
	assert.Equal(t, false, final1)
}

func TestNotExpression(t *testing.T) {
	fun, err := factory.NewExpr(`boolean.not(boolean.false())`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v1, _ := json.Marshal(fun)
	fmt.Println(string(v1))
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, true, v)
}
