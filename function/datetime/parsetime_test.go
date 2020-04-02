package datetime

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseTimes(t *testing.T) {

	fmt.Println(time.Now())
	tests := []struct {
		Date     string
		Timezone string
		Expected string
	}{
		{
			Date:     "2020-03-19T15:02:03+06:00",
			Timezone: "America/Los_Angeles",
			Expected: "2020-03-19T02:02:03-07:00",
		},
		{
			Date:     "2020-03-30T19:23:41Z",
			Expected: "2020-03-30T19:23:41Z",
		},
		{
			Date:     "September 17, 2012 at 10:09am PST-08",
			Timezone: "UTC",
			Expected: "2012-09-17T18:09:00Z",
		},
		{
			Date:     "2020-03-30T21:10:03-05:00",
			Timezone: "America/Los_Angeles",
			Expected: "2020-03-30T19:10:03-07:00",
		},
	}

	in := &fnParseTIme{}
	formatr := FormatDatetime{}
	for _, d := range tests {
		final, err := in.Eval(d.Date, d.Timezone)
		assert.Nil(t, err)
		s, _ := formatr.Eval(final, "2006-01-02T15:04:05Z07:00")
		assert.Equal(t, d.Expected, s)
	}
}
