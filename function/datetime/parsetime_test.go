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
			Date:     "Mon, 02 Jan 2020 15:04:05 MST",
			Timezone: "UTC",
			Expected: "2020-01-02T22:04:05Z",
		},
		{
			Date:     "2020-03-30T21:10:03-05:00",
			Timezone: "America/Los_Angeles",
			Expected: "2020-03-30T19:10:03-07:00",
		},
	}

	in := &fnParseTIme{}

	for _, d := range tests {
		final, err := in.Eval(d.Date, d.Timezone)
		assert.Nil(t, err)
		assert.Equal(t, d.Expected, final)
	}
}
