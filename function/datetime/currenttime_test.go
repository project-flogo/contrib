package datetime

import (
	"fmt"
	"github.com/project-flogo/core/data/expression/function"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	function.ResolveAliases()
}

func TestCurrentTime_Eval(t *testing.T) {
	n := CurrentTime{}
	date, _ := n.Eval(nil)
	assert.NotNil(t, date)
	fmt.Println(date)
}

func TestCurrentimeExpression(t *testing.T) {
	fun, err := factory.NewExpr(`datetime.currentTime()`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
