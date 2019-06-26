package datetime

import (
	"fmt"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	function.ResolveAliases()
}

func TestNow_Eval(t *testing.T) {
	n := Now{}
	now, _ := n.Eval(nil)
	assert.NotNil(t, now)
	fmt.Println(now)
}
