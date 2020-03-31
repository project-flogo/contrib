package datetime

import (
	"github.com/project-flogo/core/data"
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
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (s *fnSubHours) Eval(in ...interface{}) (interface{}, error) {
	startTime, err := ParseTime(in[0])
	if err != nil {
		return nil, err
	}
	endTime, err := ParseTime(in[1])
	if err != nil {
		return nil, err
	}

	sub := endTime.Sub(startTime).Hours()
	return sub, nil

}
