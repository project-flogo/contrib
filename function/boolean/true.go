package boolean

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
)

type True struct {
}

func init() {
	function.Register(&True{})
}

func (s *True) Name() string {
	return "true"
}

func (s *True) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

func (s *True) GetCategory() string {
	return "boolean"
}

func (s *True) Eval(d ...interface{}) (interface{}, error) {
	log.RootLogger().Debugf("Always returns true")
	return true, nil
}
