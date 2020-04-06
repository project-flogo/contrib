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

func TestFnAddSeconds_Eval(t *testing.T) {
	var in = &fnAddSeconds{}
	tests := []struct {
		Date     string
		Seconds  int
		Expected string
	}{
		{
			Date:     "2020-03-19T15:02:03Z",
			Seconds:  3,
			Expected: "2020-03-19T15:02:06Z",
		},
		{
			Date:     "2020-03-19T15:02:03",
			Seconds:  3,
			Expected: "2020-03-19T15:02:06Z",
		},
		{
			Date:     "2020-03-19T15:02:03-05:00",
			Seconds:  3,
			Expected: "2020-03-19T15:02:06-05:00",
		},
	}

	for _, d := range tests {
		final, err := in.Eval(d.Date, d.Seconds)
		assert.Nil(t, err)
		assert.Equal(t, d.Expected, final.(time.Time).Format(time.RFC3339))
	}
}
