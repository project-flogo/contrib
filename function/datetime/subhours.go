package datetime

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

type fnSubHours struct {
}

func init() {
	function.Register(&fnSubHours{})
}

func (s *fnSubHours) Name() string {
	return "subHours"
}

func (s *fnSubHours) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeDateTime, data.TypeDateTime}, false
}

func (s *fnSubHours) Eval(in ...interface{}) (interface{}, error) {
	startTime, err := coerce.ToDateTime(in[0])
	if err != nil {
		return nil, err
	}
	endTime, err := coerce.ToDateTime(in[1])
	if err != nil {
		return nil, err
	}

	sub := endTime.Sub(startTime).Hours()
	return sub, nil

}
