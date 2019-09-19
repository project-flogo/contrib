package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sp = &fnSplit{}

func TestStaticFunc_Split(t *testing.T) {
	final1, _ := sp.Eval("TIBCO Web Integrator", " ")
	final := final1.([]string)
	assert.Equal(t, "TIBCO", final[0])
	assert.Equal(t, "Web", final[1])
	assert.Equal(t, "Integrator", final[2])

	final2, _ := sp.Eval("TIBCO。网路。Integrator", "。")
	fmt.Println(final2)
	final = final2.([]string)

	assert.Equal(t, "TIBCO", final[0])
	assert.Equal(t, "网路", final[1])
	assert.Equal(t, "Integrator", final[2])

}
