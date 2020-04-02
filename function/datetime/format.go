package datetime

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

type fnFormat struct {
}

func init() {
	function.Register(&fnFormat{})
}

func (s *fnFormat) Name() string {
	return "format"
}

func (s *fnFormat) GetCategory() string {
	return "datetime"
}

func (s *fnFormat) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeDateTime, data.TypeString}, false
}

func (s *fnFormat) Eval(params ...interface{}) (interface{}, error) {
	date, err := coerce.ToDateTime(params[0])
	if err != nil {
		return nil, fmt.Errorf("Format datetime first argument must be string")
	}
	f, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("Format datetime second argument must be string")
	}
	return date.Format(format(f)), nil
}
