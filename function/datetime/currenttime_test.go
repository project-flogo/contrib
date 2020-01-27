package datetime

import (
	"testing"

	"github.com/project-flogo/core/data/expression/function"

	"github.com/stretchr/testify/assert"
)

func init() {
	function.ResolveAliases()
}

func TestCurrentTime_Eval(t *testing.T) {
	n := CurrentTime{}
	date, _ := n.Eval(nil)
	assert.NotNil(t, date)
}
