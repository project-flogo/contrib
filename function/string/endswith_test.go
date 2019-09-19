package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var end = &EndsWith{}

func TestStaticEndWith(t *testing.T) {
	final1, _ := end.Eval("TIBCO Project Flogo", "Flogo")
	fmt.Println(final1)
	assert.Equal(t, true, final1)

	final2, _ := end.Eval("TIBCO Web 集成器", "集成器")
	fmt.Println(final2)
	assert.Equal(t, true, final2)

	final3, _ := end.Eval("TIBCO 网路 FLogo", "网路")
	fmt.Println(final3)
	assert.Equal(t, false, final3)
}
