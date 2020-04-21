package datetime

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/support/log"
	"time"

	"github.com/project-flogo/core/data/expression/function"
)


type currentFn struct {
}

func init() {
	function.Register(&currentFn{})
}

func (s *currentFn) Name() string {
	return "current"
}

func (s *currentFn) GetCategory() string {
	return "datetime"
}

func (s *currentFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

func (s *currentFn) Eval(d ...interface{}) (interface{}, error) {
	log.RootLogger().Debugf("Returns the current datetime with UTC timezone")
	return time.Now().UTC(), nil
}
