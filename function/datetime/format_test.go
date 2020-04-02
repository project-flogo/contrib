package datetime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatExpression(t *testing.T) {

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

	n := &fnFormat{}

	for _, test := range testCases {
		v, err := n.Eval(test.dateV, test.format)
		assert.Nil(t, err)
		assert.NotNil(t, v)
		assert.Equal(t, test.expected, v)
	}
}

func TestFormatPredefineLayout(t *testing.T) {

	testCases := []struct {
		dateV    string
		expected string
		format   string
	}{

		{
			dateV:    "2017-04-10T22:17:32.000+0700",
			format:   "ANSIC",
			expected: "Mon Apr 10 22:17:32 2017",
		},
		{
			dateV:    "2017-04-10T22:17:32.000+0700",
			format:   "UnixDate",
			expected: "Mon Apr 10 22:17:32 +0700 2017",
		},
		{
			dateV:    "2017-04-10T22:17:32.000+0700",
			format:   "RubyDate",
			expected: "Mon Apr 10 22:17:32 +0700 2017",
		},
		{
			dateV:    "2017-04-10T22:17:32.000+0700",
			format:   "RFC822",
			expected: "10 Apr 17 22:17 +0700",
		},
		{
			dateV:    "2017-04-10T22:17:32.000+0700",
			format:   "RFC822Z",
			expected: "10 Apr 17 22:17 +0700",
		},
		{
			dateV:    "2017-04-10T22:17:32.000+0700",
			format:   "RFC850",
			expected: "Monday, 10-Apr-17 22:17:32 +0700",
		},
		{
			dateV:    "2017-04-10T22:17:32.000+0700",
			format:   "RFC1123",
			expected: "Mon, 10 Apr 2017 22:17:32 +0700",
		},
		{
			dateV:    "2017-04-10T22:17:32.000+0700",
			format:   "RFC1123Z",
			expected: "Mon, 10 Apr 2017 22:17:32 +0700",
		},
		{
			dateV:    "2017-04-10T22:17:32.000+0700",
			format:   "RFC3339",
			expected: "2017-04-10T22:17:32+07:00",
		},
		{
			dateV:    "2017-04-10T22:17:32.000+0700",
			format:   "RFC3339Nano",
			expected: "2017-04-10T22:17:32+07:00",
		},
	}

	n := &fnFormat{}

	for _, test := range testCases {
		v, err := n.Eval(test.dateV, test.format)
		assert.Nil(t, err)
		assert.NotNil(t, v)
		assert.Equal(t, test.expected, v)
	}
}
