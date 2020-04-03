package datetime

import (
	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func init() {
	function.ResolveAliases()
}

func TestFnSub_Eval(t *testing.T) {

	tests := []struct {
		Time     string
		Years    int
		Months   int
		Days     int
		Expected string
	}{
		{
			Time:     "2020-03-19T15:02:03",
			Years:    0,
			Months:   0,
			Days:     0,
			Expected: "2020-03-19T15:02:03Z",
		},
		{
			Time:     "2020-03-19T15:02:03",
			Years:    1,
			Months:   0,
			Days:     0,
			Expected: "2019-03-19T15:02:03Z",
		},
		{
			Time:     "2020-03-19T15:02:03",
			Years:    0,
			Months:   1,
			Days:     0,
			Expected: "2020-02-19T15:02:03Z",
		},
		{
			Time:     "2020-03-19T15:02:03",
			Years:    0,
			Months:   0,
			Days:     1,
			Expected: "2020-03-18T15:02:03Z",
		},
		{
			Time:     "2020-03-19T15:02:03",
			Years:    1,
			Months:   1,
			Days:     1,
			Expected: "2019-02-18T15:02:03Z",
		},
	}

	in := &fnSub{}
	for _, d := range tests {
		final, err := in.Eval(d.Time, d.Years, d.Months, d.Days)
		assert.Nil(t, err)
		assert.Equal(t, d.Expected, final.(time.Time).Format(time.RFC3339))
	}
}
