package datetime

import (
	"os"
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func init() {
	function.ResolveAliases()
}

func TestCurrentDatetime_Eval(t *testing.T) {
	n := CurrentDatetime{}
	datetime, _ := n.Eval(nil)
	assert.NotNil(t, datetime)
}

func TestDatetime_CDT(t *testing.T) {
	os.Setenv(WI_DATETIME_LOCATION, "US/Central")
	n := CurrentDatetime{}
	date, _ := n.Eval(nil)
	assert.NotNil(t, date)
}
