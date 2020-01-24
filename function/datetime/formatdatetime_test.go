package datetime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatDatetime_Eval(t *testing.T) {
	n := &FormatDatetime{}

	date, err := n.Eval("2017-04-12T22:15:09", "2006-01-02 15:04:05")
	assert.Nil(t, err)
	assert.NotNil(t, date)
	assert.Equal(t, "2017-04-12 22:15:09", date)

}

func TestFormatDatetime_Eval2(t *testing.T) {
	n := &FormatDatetime{}
	date, err := n.Eval("2017-04-10T22:17:32.000+0000", "2006-01-02 15:04:05")
	assert.Nil(t, err)
	assert.NotNil(t, date)
	assert.Equal(t, "2017-04-10 22:17:32", date)

}

//func TestFormatDatetime_Default(t *testing.T) {
//	n := &FormatDatetime{}
//	date, err := n.Eval("2017-04-10T22:17:32.000+0000", GetDatetimeFormat())
//	assert.Nil(t, err)
//	assert.NotNil(t, date)
//	assert.Equal(t, "2017-04-10T22:17:32+00:00", date)
//	fmt.Println(date)
//}

func TestFormatDateTimeExpression(t *testing.T) {

	testCases := []struct {
		dateV    string
		expected string
		format   string
	}{

		{
			dateV:    "2017-04-10T22:17:32.000+0000",
			format:   "2006-01-02 15:04:05",
			expected: "2017-04-10 22:17:32",
		},
		{
			dateV:    "2017-04-10T22:17:32.000+0700",
			format:   "dd-MM-yyyy T hh-mm-ss",
			expected: "10-04-2017 T 22-17-32",
		},
		{
			dateV:    "2017-04-10T22:17:32.000+0000",
			format:   "dd/MM/yyyy T hh:mm:ss",
			expected: "10/04/2017 T 22:17:32",
		},
		{
			dateV:    "2017-04-10T22:17:32.000+0000",
			format:   "2006-01-02 15:04:05",
			expected: "2017-04-10 22:17:32",
		},
	}

	n := &FormatDatetime{}

	for _, test := range testCases {
		v, err := n.Eval(test.dateV, test.format)
		assert.Nil(t, err)
		assert.NotNil(t, v)
		assert.Equal(t, test.expected, v)
	}
}
