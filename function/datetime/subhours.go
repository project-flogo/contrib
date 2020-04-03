package datetime

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"strconv"
	"time"
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
	return []data.Type{data.TypeDateTime, data.TypeFloat64}, false
}

func (s *fnSubHours) Eval(in ...interface{}) (interface{}, error) {
	t, err := coerce.ToDateTime(in[0])
	if err != nil {
		return nil, err
	}
	hours, err := coerce.ToFloat64(in[1])
	if err != nil {
		return nil, err
	}
	d, err := time.ParseDuration("-" + strconv.FormatFloat(hours, 'f', -1, 64) + "h")
	if err != nil {
		return nil, fmt.Errorf("Invalid hours [%f]", hours)
	}
	return t.Add(d), nil
}
