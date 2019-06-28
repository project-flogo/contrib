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

type FormatDatetime struct {
}

func init() {
	function.Register(&FormatDatetime{})
}

func (s *FormatDatetime) Name() string {
	return "formatDatetime"
}

func (s *FormatDatetime) GetCategory() string {
	return "datetime"
}

func (s *FormatDatetime) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (s *FormatDatetime) Eval(params ...interface{}) (interface{}, error) {
	date, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("Format datetime first argument must be string")
	}
	format, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("Format datetime second argument must be string")
	}

	log.RootLogger().Debugf("Format datetime %s to format %s", date, format)
	p, err := parsetime.NewParseTime(GetLocation())
	if err != nil {
		log.RootLogger().Errorf("New datetime parser %s error %s", date, err.Error())
		return date, err
	}
	t, err := p.Parse(date)
	if err != nil {
		log.RootLogger().Errorf("Parsing datetime %s error %s", date, err.Error())
		return date, err
	}
	return t.Format(convertDateTimeFormater(format)), nil
}

func convertDateTimeFormater(format string) string {

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
