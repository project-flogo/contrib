package datetime

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
	"strings"
)

// Deprecated
type FormatDate struct {
}

func init() {
	function.Register(&FormatDate{})
}

func (s *FormatDate) Name() string {
	return "formatDate"
}

func (s *FormatDate) GetCategory() string {
	return "datetime"
}

func (s *FormatDate) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeDateTime, data.TypeString}, false
}

func (s *FormatDate) Eval(params ...interface{}) (interface{}, error) {
	date, err := coerce.ToDateTime(params[0])
	if err != nil {
		return nil, fmt.Errorf("Format date first argument must be string")
	}
	format, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("Format date second argument must be string")
	}

	format = convertDateFormater(format)
	log.RootLogger().Debugf("Format date %s to format %s", date, format)
	return date.Format(format), nil
}

func convertDateFormater(format string) string {
	lowerFormat := strings.ToLower(format)

	if strings.Contains(lowerFormat, "yyyy") {
		lowerFormat = strings.Replace(lowerFormat, "yyyy", "2006", -1)
	}

	if strings.Contains(lowerFormat, "mm") {
		lowerFormat = strings.Replace(lowerFormat, "mm", "01", -1)
	}

	if strings.Contains(lowerFormat, "dd") {
		lowerFormat = strings.Replace(lowerFormat, "dd", "02", -1)
	}
	return lowerFormat
}
