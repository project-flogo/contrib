package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var index = &fnIndex{}

func TestStatic_Index(t *testing.T) {
	final1, _ := index.Eval("TIBCO Web Integrator", "Web")
	assert.Equal(t, 6, final1)

	final2, _ := index.Eval("TIBCO Web Integrator", "Internet")
	assert.Equal(t, -1, final2)

	final3, _ := index.Eval("TIBCO 网络 Integrator", "Integrator")
	assert.Equal(t, 13, final3)
}
