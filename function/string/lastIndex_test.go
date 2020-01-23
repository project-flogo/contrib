package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var in = &fnLastIndex{}

func TestStaticFunc_Index(t *testing.T) {
	final1, _ := in.Eval("Integration with TIBCO Web Integrator", "Integrat")
	assert.Equal(t, 27, final1)

	final2, _ := in.Eval("TIBCO Web Integrator", "rocks")
	assert.Equal(t, -1, final2)

	final3, _ := in.Eval("Integration with TIBCO 网络 Integrator", "Integrat")
	assert.Equal(t, 30, final3)
}
