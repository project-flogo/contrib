package datetime

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"time"
)

type fnAddSeconds struct {
}

func init() {
	function.Register(&fnAddSeconds{})
}

func (s *fnAddSeconds) Name() string {
	return "addSeconds"
}

func (s *fnAddSeconds) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeDateTime, data.TypeInt}, false
}

func (s *fnAddSeconds) Eval(in ...interface{}) (interface{}, error) {

	startDate, err := coerce.ToDateTime(in[0])
	if err != nil {
		return nil, err
	}

	seconds, err := coerce.ToInt64(in[1])
	if err != nil {
		return nil, err
	}

	newT := startDate.Add(time.Duration(seconds) * time.Second)
	return newT, nil

}
