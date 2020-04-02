package datetime

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/support/log"
	"time"

	"github.com/project-flogo/core/data/expression/function"
)

const DateFormatDefault = "2006-01-02-07:00"

type CurrentDate struct {
}

func init() {
	function.Register(&CurrentDate{})
}

func (s *CurrentDate) Name() string {
	return "currentDate"
}

func (s *CurrentDate) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

func (s *CurrentDate) Eval(d ...interface{}) (interface{}, error) {
	log.RootLogger().Debugf("Returns the current date with timezone")
	return time.Now().UTC().Format(DateFormatDefault), nil
}
