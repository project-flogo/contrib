package datetime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFnDiff_Eval(t *testing.T) {

	tests := []struct {
		Time     string
		Time2    string
		Return   string
		Expected float64
	}{
		{
			Time:     "2020-03-19T15:02:03",
			Time2:    "2020-03-22T15:02:03",
			Return:   "days",
			Expected: float64(3),
		},
		{
			Time:     "2020-03-19T15:02:03",
			Time2:    "2020-03-22T12:05:03",
			Return:   "days",
			Expected: float64(2.877083333333333),
		},
		{
			Time:   "2020-03-19T15:02:03",
			Time2:  "2020-03-22T15:02:03",
			Return: "hours",

			Expected: float64(72),
		},
		{
			Time:     "2020-03-19T15:02:03",
			Time2:    "2020-03-22T12:05:03",
			Return:   "hours",
			Expected: float64(69.05),
		},
		{
			Time:     "2020-03-19T15:02:03",
			Time2:    "2020-03-22T15:02:03",
			Return:   "mins",
			Expected: float64(4320),
		},
		{
			Time:     "2020-03-19T15:02:03",
			Time2:    "2020-03-19T15:05:03",
			Return:   "mins",
			Expected: float64(3),
		},
		{
			Time:     "2020-03-19T15:02:03",
			Time2:    "2020-03-19T15:07:22",
			Return:   "seconds",
			Expected: float64(319),
		},
		{
			Time:     "2020-03-19T15:02:03",
			Time2:    "2020-03-22T12:05:03",
			Return:   "seconds",
			Expected: float64(248580),
		},
	}

	in := &fnDiff{}
	for _, d := range tests {
		final, err := in.Eval(d.Time, d.Time2, d.Return)
		assert.Nil(t, err)
		assert.Equal(t, d.Expected, final)
	}
}
