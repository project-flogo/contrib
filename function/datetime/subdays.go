package datetime

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
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
	return []data.Type{data.TypeDateTime, data.TypeDateTime}, false
}

func (s *fnSubDays) Eval(in ...interface{}) (interface{}, error) {

	startTime, err := coerce.ToDateTime(in[0])
	if err != nil {
		return nil, err
	}
	endTime, err := coerce.ToDateTime(in[1])
	if err != nil {
		return nil, err
	}

	sub := endTime.Sub(startTime).Hours() / 24

	return sub, nil

}
