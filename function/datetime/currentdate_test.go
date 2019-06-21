package datetime

import (
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/data/expression/script"
	"github.com/project-flogo/core/data/resolve"
	"testing"

	"os"

	"fmt"
	"github.com/stretchr/testify/assert"
)

var resolver = resolve.NewCompositeResolver(map[string]resolve.Resolver{"static": nil, ".": nil, "env": &resolve.EnvResolver{}})
var factory = script.NewExprFactory(resolver)

func init() {
	function.ResolveAliases()
}

func TestCurrectDaye_Eval(t *testing.T) {
	n := CurrentDate{}
	date, _ := n.Eval(nil)
	assert.NotNil(t, date)
}

func TestNow_CDT(t *testing.T) {
	os.Setenv(WI_DATETIME_LOCATION, "US/Central")
	n := CurrentDate{}
	date, _ := n.Eval(nil)
	assert.NotNil(t, date)
}

func TestCurrectDateExpression(t *testing.T) {
	fun, err := factory.NewExpr(`datetime.currentDate()`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
