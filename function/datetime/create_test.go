package datetime

import (
	"fmt"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func init() {
	function.ResolveAliases()
}

func TestFnCreate_Eval(t *testing.T) {
	var in = &fnCreate{}
	tests := []struct {
		Years    int
		Months   int
		Days     int
		HH       int
		MM       int
		SS       int
		NS       int
		Loc      string
		Expected string
	}{
		{
			Years:    2020,
			Months:   1,
			Days:     3,
			HH:       2,
			MM:       22,
			Loc:      "America/Los_Angeles",
			Expected: "2020-01-03T02:22:00-08:00",
		},
		{
			Years:    2020,
			Months:   1,
			Days:     3,
			Loc:      "America/Los_Angeles",
			Expected: "2020-01-03T00:00:00-08:00",
		},
	}

	for _, d := range tests {
		final, err := in.Eval(d.Years, d.Months, d.Days, d.HH, d.MM, d.SS, d.NS, d.Loc)
		assert.Nil(t, err)
		fmt.Println(final)
		assert.Equal(t, d.Expected, FormatDateWithRFC3339(final.(time.Time)))
	}
}
