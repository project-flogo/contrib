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

func TestFnSubMins_Eval(t *testing.T) {

	tests := []struct {
		Time     string
		Mins     int
		Expected string
	}{
		{
			Time:     "2020-03-19T15:02:03",
			Mins:     10,
			Expected: "2020-03-19T14:52:03Z",
		},
		{
			Time:     "2020-03-19T15:02:03",
			Mins:     30,
			Expected: "2020-03-19T14:32:03Z",
		},
	}

	in := &fnSubMins{}

	for _, d := range tests {
		final, err := in.Eval(d.Time, d.Mins)
		assert.Nil(t, err)
		assert.Equal(t, d.Expected, final.(time.Time).Format(time.RFC3339))
	}
}
