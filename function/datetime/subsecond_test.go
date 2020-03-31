package datetime

import (
	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	function.ResolveAliases()
}

func TestFnSubSecond_Eval(t *testing.T) {

	tests := []struct {
		Time     string
		Time2    string
		Expected float64
	}{
		{
			Time:     "2020-03-19T15:02:03",
			Time2:    "2020-03-19T15:07:22",
			Expected: float64(319),
		},
		{
			Time:     "2020-03-19T15:02:03",
			Time2:    "2020-03-22T12:05:03",
			Expected: float64(248580),
		},
	}

	in := &fnSubSecond{}

	for _, d := range tests {
		final, err := in.Eval(d.Time, d.Time2)
		assert.Nil(t, err)
		assert.Equal(t, d.Expected, final)
	}
}
