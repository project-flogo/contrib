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
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (s *FormatTime) Eval(params ...interface{}) (interface{}, error) {
	date, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("Format time first argument must be string")
	}
	format, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("Format time second argument must be string")
	}

	log.RootLogger().Debugf("Format time %s to format %s", date, format)
	p, err := parsetime.NewParseTime(GetLocation())
	if err != nil {
		log.RootLogger().Errorf("New time parser %s error %s", date, err.Error())
		return date, err
	}
	t, err := p.Parse(date)
	if err != nil {
		log.RootLogger().Errorf("Parsing time %s error %s", date, err.Error())
		return date, err
	}
	return t.Format(convertTimeFormater(format)), nil
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
