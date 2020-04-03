package datetime

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"strconv"
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
	return []data.Type{data.TypeDateTime, data.TypeFloat64}, false
}

func (s *fnAddHours) Eval(in ...interface{}) (interface{}, error) {
	startDate, err := coerce.ToDateTime(in[0])
	if err != nil {
		return nil, err
	}
	hours, err := coerce.ToFloat64(in[1])
	if err != nil {
		return nil, err
	}

	d, err := time.ParseDuration(strconv.FormatFloat(hours, 'f', -1, 64) + "h")
	if err != nil {
		return nil, fmt.Errorf("Invalid hours [%f]", hours)
	}
	newT := startDate.Add(d)
	return newT, nil
}
