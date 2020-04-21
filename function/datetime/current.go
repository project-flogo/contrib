package datetime

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/support/log"
	"time"

	"github.com/project-flogo/core/data/expression/function"
)


type curremtFn struct {
}

func init() {
	function.Register(&curremtFn{})
}

func (s *curremtFn) Name() string {
	return "current"
}

func (s *curremtFn) GetCategory() string {
	return "datetime"
}

func (s *curremtFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

func (s *curremtFn) Eval(d ...interface{}) (interface{}, error) {
	log.RootLogger().Debugf("Returns the current datetime with UTC timezone")
	return time.Now().UTC(), nil
}
