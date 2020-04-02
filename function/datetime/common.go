package datetime

import "time"

func FormatDateWithRFC3339(t time.Time) string {
	return t.Format(time.RFC3339)
}
