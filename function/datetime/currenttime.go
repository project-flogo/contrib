package datetime

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/support/log"
	"time"

	"github.com/project-flogo/core/data/expression/function"
)

const TimeFormatDefault string = "15:04:05-07:00"

type CurrentTime struct {
}

func init() {
	function.Register(&CurrentTime{})
}

func (s *CurrentTime) Name() string {
	return "currentTime"
}

func (s *CurrentTime) GetCategory() string {
	return "datetime"
}

func (s *CurrentTime) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

func (s *CurrentTime) Eval(d ...interface{}) (interface{}, error) {
	log.RootLogger().Debugf("Returns the current time with timezone")
	return time.Now().UTC().Format(TimeFormatDefault), nil
}
