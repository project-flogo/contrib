package string

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var index = &fnIndex{}

func TestStatic_Index(t *testing.T) {
	final1, _ := index.Eval("TIBCO Web Integrator", "Web")
	fmt.Println(final1)
	assert.Equal(t, 6, final1)

	final2, _ := index.Eval("TIBCO Web Integrator", "Internet")
	fmt.Println(final2)
	assert.Equal(t, -1, final2)

	final3, _ := index.Eval("TIBCO 网络 Integrator", "Integrator")
	fmt.Println(final3)
	assert.Equal(t, 13, final3)
}
