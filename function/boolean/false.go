package boolean

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
)

type False struct {
}

func init() {
	function.Register(&False{})
}

func (s *False) Name() string {
	return "false"
}

func (s *False) GetCategory() string {
	return "boolean"
}

func (s *False) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

func (s *False) Eval(data ...interface{}) (interface{}, error) {
	log.RootLogger().Debugf("Always returns false")
	return false, nil
}
