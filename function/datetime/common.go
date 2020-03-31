package datetime

import (
	"fmt"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/support/log"
	"github.com/tkuchiki/parsetime"
	"time"
)

func ParseTime(t interface{}) (time.Time, error) {
	var parsedTime time.Time
	switch tt := t.(type) {
	case time.Time:
		parsedTime = tt
	default:
		date, err := coerce.ToString(t)
		if err != nil {
			return parsedTime, fmt.Errorf("Format date first argument must be string")
		}

		p, err := parsetime.NewParseTime(GetLocation())
		if err != nil {
			log.RootLogger().Errorf("New date parser %s error %s", date, err.Error())
			return parsedTime, err
		}

		parsedTime, err = p.Parse(date)
		if err != nil {
			log.RootLogger().Errorf("Parsing date %s error %s", date, err.Error())
			return parsedTime, err
		}
	}
	return parsedTime, nil
}
