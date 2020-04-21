package datetime

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func init() {
	function.ResolveAliases()
}

func TestCurrentDate_Eval(t *testing.T) {
	n := currentFn{}
	datetime, _ := n.Eval(nil)
	fmt.Println(datetime)
	assert.NotNil(t, datetime)
}

