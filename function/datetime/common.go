package datetime

import (
	"strings"
	"time"
)

func FormatDateWithRFC3339(t time.Time) string {
	return t.Format(format("RFC3339"))
}

func format(format string) string {

	/*
		    ANSIC       = "Mon Jan _2 15:04:05 2006"
			UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
			RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
			RFC822      = "02 Jan 06 15:04 MST"
			RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
			RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
			RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
			RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
			RFC3339     = "2006-01-02T15:04:05Z07:00"
			RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	*/

	if strings.EqualFold(format, "ANSIC") {
		return time.ANSIC
	} else if strings.EqualFold(format, "UnixDate") {
		return time.UnixDate
	} else if strings.EqualFold(format, "RubyDate") {
		return time.RubyDate
	} else if strings.EqualFold(format, "RFC822") {
		return time.RFC822
	} else if strings.EqualFold(format, "RFC822Z") {
		return time.RFC822Z
	} else if strings.EqualFold(format, "RFC850") {
		return time.RFC850
	} else if strings.EqualFold(format, "RFC1123") {
		return time.RFC1123
	} else if strings.EqualFold(format, "RFC1123Z") {
		return time.RFC1123Z
	} else if strings.EqualFold(format, "RFC3339") {
		return time.RFC3339
	} else if strings.EqualFold(format, "RFC3339Nano") {
		return time.RFC3339Nano
	}

	if strings.Contains(strings.ToLower(format), "yyyy") {
		format = strings.Replace(format, "yyyy", "2006", -1)
		format = strings.Replace(format, "YYYY", "2006", -1)
	}

	if strings.Contains(format, "MM") {
		format = strings.Replace(format, "MM", "01", -1)
	}

	if strings.Contains(strings.ToLower(format), "dd") {
		format = strings.Replace(format, "dd", "02", -1)
		format = strings.Replace(format, "DD", "02", -1)

	}

	if strings.Contains(strings.ToLower(format), "hh") {
		format = strings.Replace(format, "hh", "15", -1)
		format = strings.Replace(format, "HH", "15", -1)
	}

	if strings.Contains(format, "mm") {
		format = strings.Replace(format, "mm", "04", -1)
	}

	if strings.Contains(strings.ToLower(format), "ss") {
		format = strings.Replace(format, "ss", "05", -1)
		format = strings.Replace(format, "SS", "05", -1)

	}
	return format
}
