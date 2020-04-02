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

func TestFnAdd_Eval(t *testing.T) {
	var in = &fnAdd{}
	tests := []struct {
		Date     string
		Years    int
		Months   int
		Days     int
		Expected string
	}{
		{
			Date:     "2020-03-19T15:02:03Z",
			Years:    1,
			Months:   1,
			Days:     3,
			Expected: "2021-04-22T15:02:03Z",
		},
		{
			Date:     "2020-03-19T15:02:03",
			Years:    0,
			Months:   0,
			Days:     3,
			Expected: "2020-03-22T15:02:03Z",
		},
		{
			Date:     "2020-03-19T15:02:03-05:00",
			Years:    0,
			Months:   2,
			Days:     0,
			Expected: "2020-05-19T15:02:03-05:00",
		},
	}

	for _, d := range tests {
		final, err := in.Eval(d.Date, d.Years, d.Months, d.Days)
		assert.Nil(t, err)
		assert.Equal(t, d.Expected, final.(time.Time).Format(time.RFC3339))
	}
}
