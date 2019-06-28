package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var start = &StartsWith{}

func TestStaticFunc_Starts_with(t *testing.T) {
	final1, _ := start.Eval("TIBCO Web Integrator", "TIBCO")
	fmt.Println(final1)
	assert.Equal(t, true, final1)

	final2, _ := start.Eval("网路 Integrator", "网路")
	fmt.Println(final2)
	assert.Equal(t, true, final2)

	final3, _ := start.Eval("TIBCO 网路 Integrator", "网路")
	fmt.Println(final3)
	assert.Equal(t, false, final3)
}
