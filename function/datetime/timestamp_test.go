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

func TestGetTimestamp_Eval(t *testing.T) {
	n := Timestamp{}
	timestamp, err := n.Eval(nil)
	fmt.Println(timestamp)
	assert.Nil(t, err)
	assert.NotNil(t, timestamp)

}
