package datetime

import (
	"fmt"
	"github.com/project-flogo/core/data/expression/function"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func init() {
	function.ResolveAliases()
}

func TestCurrentDatetime_Eval(t *testing.T) {
	n := CurrentDatetime{}
	datetime, _ := n.Eval(nil)
	assert.NotNil(t, datetime)
	logrus.Info(datetime)
	fmt.Println(datetime)
}

func TestDatetime_CDT(t *testing.T) {
	os.Setenv(WI_DATETIME_LOCATION, "US/Central")
	n := CurrentDatetime{}
	date, _ := n.Eval(nil)
	assert.NotNil(t, date)
	logrus.Info(date)
}
