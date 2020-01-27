package datetime

import (
	"testing"

	"github.com/project-flogo/core/data/expression/function"

	"github.com/stretchr/testify/assert"
)

func init() {
	function.ResolveAliases()
}

func TestFormatDate_Eval(t *testing.T) {
	n := FormatDate{}

	date, err := n.Eval("02/08/2017", "20060102")
	assert.Nil(t, err)
	assert.NotNil(t, date)
	assert.Equal(t, "20170208", date)
}

func TestFormatDate_Eval2(t *testing.T) {
	n := FormatDate{}
	date, err := n.Eval("02/08/2017", "2006-02-01")
	assert.Nil(t, err)
	assert.NotNil(t, date)
	assert.Equal(t, "2017-08-02", date)
}

func TestFormatDate_Eval3(t *testing.T) {
	n := FormatDate{}
	date, err := n.Eval("02/08/2017", "01-02-2006")
	assert.Nil(t, err)
	assert.NotNil(t, date)
	assert.Equal(t, "02-08-2017", date)
}

func TestFormatDateYYYYDDDD(t *testing.T) {

	testCases := []struct {
		dateV    string
		expected string
		format   string
	}{

		{
			dateV:    "02/08/2017",
			format:   "yyyymmdd",
			expected: "20170208",
		},
		{
			dateV:    "02/08/2017",
			format:   "ddMMyyyy",
			expected: "08022017",
		},
		{
			dateV:    "02/08/2017",
			format:   "dd-MM-yyyy",
			expected: "08-02-2017",
		},
	}

	for _, test := range testCases {
		n := FormatDate{}
		date, err := n.Eval(test.dateV, test.format)
		assert.Nil(t, err)
		assert.NotNil(t, date)
		assert.Equal(t, test.expected, date)
	}

}
