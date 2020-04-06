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
	return []data.Type{data.TypeDateTime, data.TypeInt}, false
}

func (s *fnSubSecond) Eval(in ...interface{}) (interface{}, error) {
	t, err := coerce.ToDateTime(in[0])
	if err != nil {
		return nil, err
	}
	seconds, err := coerce.ToInt(in[1])
	if err != nil {
		return nil, err
	}

	d, err := time.ParseDuration("-" + strconv.Itoa(seconds) + "s")
	if err != nil {
		return nil, fmt.Errorf("Invalid minutes [%d]", seconds)
	}
	return t.Add(d), nil
}
