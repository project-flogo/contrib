package string

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var in = &fnLastIndex{}

func TestStaticFunc_Index(t *testing.T) {
	final1, _ := in.Eval("Integration with TIBCO Web Integrator", "Integrat")
	fmt.Println(final1)
	assert.Equal(t, 27, final1)

	final2, _ := in.Eval("TIBCO Web Integrator", "rocks")
	fmt.Println(final2)
	assert.Equal(t, -1, final2)

	final3, _ := in.Eval("Integration with TIBCO 网络 Integrator", "Integrat")
	fmt.Println(final3)
	assert.Equal(t, 30, final3)
}
