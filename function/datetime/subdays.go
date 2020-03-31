package datetime

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

type fnSubDays struct {
}

func init() {
	function.Register(&fnSubDays{})
}

func (s *fnSubDays) Name() string {
	return "subDays"
}

func (s *fnSubDays) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (s *fnSubDays) Eval(in ...interface{}) (interface{}, error) {

	startTime, err := ParseTime(in[0])
	if err != nil {
		return nil, err
	}
	endTime, err := ParseTime(in[1])
	if err != nil {
		return nil, err
	}

	sub := endTime.Sub(startTime).Hours() / 24

	return sub, nil

}
