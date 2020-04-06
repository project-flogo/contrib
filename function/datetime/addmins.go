package datetime

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"time"
)

type fnAddMins struct {
}

func init() {
	function.Register(&fnAddMins{})
}

func (s *fnAddMins) Name() string {
	return "addMins"
}

func (s *fnAddMins) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeDateTime, data.TypeInt}, false
}

func (s *fnAddMins) Eval(in ...interface{}) (interface{}, error) {
	t, err := coerce.ToDateTime(in[0])
	if err != nil {
		return nil, err
	}
	mins, err := coerce.ToInt(in[1])
	if err != nil {
		return nil, err
	}

	return t.Add(time.Duration(mins) * time.Minute), nil

}
