package datetime

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"strconv"
	"time"
)

type fnSubSecond struct {
}

func init() {
	function.Register(&fnSubSecond{})
}

func (s *fnSubSecond) Name() string {
	return "subSeconds"
}

func (s *fnSubSecond) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeDateTime, data.TypeFloat64}, false
}

func (s *fnSubSecond) Eval(in ...interface{}) (interface{}, error) {
	t, err := coerce.ToDateTime(in[0])
	if err != nil {
		return nil, err
	}
	seconds, err := coerce.ToFloat64(in[1])
	if err != nil {
		return nil, err
	}

	d, err := time.ParseDuration("-" + strconv.FormatFloat(seconds, 'f', -1, 64) + "s")
	if err != nil {
		return nil, fmt.Errorf("Invalid minutes [%f]", seconds)
	}
	return t.Add(d), nil
}
