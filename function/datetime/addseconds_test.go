package datetime

import (
	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	function.ResolveAliases()
}

func TestFnAddSeconds_Eval(t *testing.T) {
	var in = &fnAddSeconds{}
	tests := []struct {
		Date     string
		Days     int
		Expected string
	}{
		{
			Date:     "2020-03-19T15:02:03Z",
			Days:     3,
			Expected: "2020-03-19T15:02:06Z",
		},
		{
			Date:     "2020-03-19T15:02:03",
			Days:     3,
			Expected: "2020-03-19T15:02:06Z",
		},
		{
			Date:     "2020-03-19T15:02:03-05:00",
			Days:     3,
			Expected: "2020-03-19T15:02:06-05:00",
		},
	}

	for _, d := range tests {
		final, err := in.Eval(d.Date, d.Days)
		assert.Nil(t, err)
		assert.Equal(t, d.Expected, final)
	}
}
