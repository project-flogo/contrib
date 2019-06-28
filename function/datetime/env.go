package datetime

import (
	"os"
)

const (
	WI_DATETIME_LOCATION         string = "FLOGO_DATETIME_LOCATION"
	WI_DATETIME_LOCATION_DEFAULT string = "UTC"
)

func GetLocation() string {
	location, ok := os.LookupEnv(WI_DATETIME_LOCATION)
	if ok && location != "" {
		return location
	}
	return WI_DATETIME_LOCATION_DEFAULT
}
