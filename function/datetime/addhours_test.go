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

func TestFnAddHours_Eval(t *testing.T) {
	var in = &fnAddHours{}
	tests := []struct {
		Date     string
		Hours    int
		Expected string
	}{
		{
			Date:     "2020-03-19T15:02:03Z",
			Hours:    3,
			Expected: "2020-03-19T18:02:03Z",
		},
		{
			Date:     "2020-03-19T15:02:03",
			Hours:    3,
			Expected: "2020-03-19T18:02:03Z",
		},
		{
			Date:     "2020-03-19T15:02:03-05:00",
			Hours:    3,
			Expected: "2020-03-19T18:02:03-05:00",
		},
	}

	for _, d := range tests {
		final, err := in.Eval(d.Date, d.Hours)
		assert.Nil(t, err)
		assert.Equal(t, d.Expected, final.(time.Time).Format(time.RFC3339))
	}
}
