package boolean

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
)


type Not struct {
}

func init() {
	function.Register(&Not{})
}

func (s *Not) Name() string {
	return "not"
}

func (s *Not) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeBool}, false
}

func (s *Not) GetCategory() string {
	return "boolean"
}

func (s *Not) Eval(values ...interface{}) (interface{}, error) {
	value := values[0].(bool)
	log.RootLogger().Debugf("Not %v", value)
	return !value, nil
}
