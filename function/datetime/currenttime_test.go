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
