package datetime

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/support/log"
	"time"

	"github.com/project-flogo/core/data/expression/function"
)

type Now struct {
}

func init() {
	function.Register(&Now{})
}

func (s *Now) Name() string {
	return "now"
}

func (s *Now) GetCategory() string {
	return "datetime"
}

func (s *Now) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

func (s *Now) Eval(params ...interface{}) (interface{}, error) {
	log.RootLogger().Debugf("Returns the current datetime with timezone")
	var currentTime time.Time
	location, err := time.LoadLocation(GetLocation())
	if err != nil {
		log.RootLogger().Errorf("Load location %s error %s", GetLocation(), err.Error())
		location = time.UTC
		currentTime = time.Now().UTC()
	} else {
		currentTime = time.Now().In(location)
	}
	return currentTime.Format(DateTimeFormatDefault), nil
}
