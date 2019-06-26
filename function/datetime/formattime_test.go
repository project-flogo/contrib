package datetime

import (
	"github.com/project-flogo/core/data/expression/function"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	function.ResolveAliases()
}

//func TestFormatTime_Eval(t *testing.T) {
//	n := FormatTime{}
//	date, err := n.Eval("10:11:05.00000 ", GetTimeFormat())
//	assert.Nil(t, err)
//	assert.NotNil(t, date)
//	assert.Equal(t, "10:11:05+00:00", date)
//	fmt.Println(date)
//}

func TestFormatTimeExpression(t *testing.T) {

	testCases := []struct {
		dateV    string
		expected string
		format   string
	}{

		{
			dateV:    "10:11:05.00000",
			format:   "15:04:05",
			expected: "10:11:05",
		},
		{
			dateV:    "10:11:05.00000",
			format:   "hh-mm-ss",
			expected: "10-11-05",
		},
	}
	tf := &FormatTime{}
	for _, test := range testCases {
		v, err := tf.Eval(test.dateV, test.format)
		assert.Nil(t, err)
		assert.NotNil(t, v)
		assert.Equal(t, test.expected, v)
	}
}
