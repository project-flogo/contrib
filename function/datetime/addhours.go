package datetime

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"time"
)

type fnAddHours struct {
}

func init() {
	function.Register(&fnAddHours{})
}

func (s *fnAddHours) Name() string {
	return "addHours"
}

func (s *fnAddHours) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeInt}, false
}

func (s *fnAddHours) Eval(in ...interface{}) (interface{}, error) {
	startDate, err := ParseTime(in[0])
	if err != nil {
		return nil, err
	}
	hours, err := coerce.ToInt64(in[1])
	if err != nil {
		return nil, err
	}
	newT := startDate.Add(time.Duration(hours) * time.Hour)
	return newT.Format(time.RFC3339), nil
}
