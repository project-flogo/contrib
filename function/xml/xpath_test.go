package xml

import (
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

var (
	expectedResult = `<transaction id="4"><name>Dave3</name><amount>400</amount></transaction>`
	xpath          = `//transactions/transaction[@id='4']`
	xml            = `<root>
    <transactions>
        <transaction id="1">
            <name>Dave</name>
            <amount>100</amount>
        </transaction>
        <transaction id="2">
            <name>Dave1</name>
            <amount>200</amount>
        </transaction>
        <transaction id="3">
            <name>Dave2</name>
            <amount>300</amount>
        </transaction>
        <transaction id="4">
            <name>Dave3</name>
            <amount>400</amount>
        </transaction>
        <transaction id="5">
            <name>Dave4</name>
            <amount>500</amount>
        </transaction>
    </transactions>
</root>`
)

func init() {
	function.ResolveAliases()
}

func TestFnxpath_Eval(t *testing.T) {
	f := &fnXPATH{}
	v, err := function.Eval(f, xpath, xml, true)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, v)
}
