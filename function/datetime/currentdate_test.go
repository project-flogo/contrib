package datetime

import (
	"github.com/project-flogo/core/data/expression/function"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	function.ResolveAliases()
}

func TestCurrectDaye_Eval(t *testing.T) {
	n := CurrentDate{}
	date, _ := n.Eval(nil)
	assert.NotNil(t, date)
}

func TestNow_CDT(t *testing.T) {
	n := CurrentDate{}
	date, _ := n.Eval(nil)
	assert.NotNil(t, date)
}
