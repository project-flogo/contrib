package datetime

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
	"github.com/tkuchiki/parsetime"
	"strings"
)

// Deprecated
type FormatTime struct {
}

func init() {
	function.Register(&FormatTime{})
}

func (s *FormatTime) Name() string {
	return "formatTime"
}

func (s *FormatTime) GetCategory() string {
	return "datetime"
}

func (s *FormatTime) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeString}, false
}

func (s *FormatTime) Eval(params ...interface{}) (interface{}, error) {
	date, err := coerce.ToDateTime(params[0])
	if err != nil {
		//For backward compatible.
		p, err := parsetime.NewParseTime("UTC")
		if err != nil {
			return nil, fmt.Errorf("parse time [%s] error: %s", params[0], err.Error())
		}

		dStr, _ := coerce.ToString(params[0])
		date, err = p.Parse(dStr)
		if err != nil {
			return nil, fmt.Errorf("parse time [%s] error: %s", params[0], err.Error())
		}
	}
	format, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("Format time second argument must be string")
	}

	log.RootLogger().Debugf("Format time %s to format %s", date, format)
	return date.Format(convertTimeFormater(format)), nil
}

//var formater = map[string]string{"DD": "02", "MM": "01", "YYYY": "2006", "HH": "15", "mm": "04", "ss": "05"}

func convertTimeFormater(format string) string {

	lowerFormat := strings.ToLower(format)

	if strings.Contains(lowerFormat, "hh") {
		lowerFormat = strings.Replace(lowerFormat, "hh", "15", -1)
	}

	if strings.Contains(lowerFormat, "mm") {
		lowerFormat = strings.Replace(lowerFormat, "mm", "04", -1)
	}

	if strings.Contains(lowerFormat, "ss") {
		lowerFormat = strings.Replace(lowerFormat, "ss", "05", -1)
	}
	return lowerFormat
}
