package datetime

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"strings"
)

const (
	Days    = "days"
	Hours   = "hours"
	Mins    = "mins"
	Seconds = "seconds"
)

type fnDiff struct {
}

func init() {
	function.Register(&fnDiff{})
}

func (s *fnDiff) Name() string {
	return "diff"
}

func (s *fnDiff) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeDateTime, data.TypeDateTime, data.TypeString}, false
}

func (s *fnDiff) Eval(in ...interface{}) (interface{}, error) {
	start, err := coerce.ToDateTime(in[0])
	if err != nil {
		return nil, err
	}
	end, err := coerce.ToDateTime(in[1])
	if err != nil {
		return nil, err
	}

	returnType, err := coerce.ToString(in[2])
	if err != nil {
		return nil, err
	}

	diff := end.Sub(start)

	if strings.EqualFold(returnType, Days) {
		return diff.Hours() / 24, nil
	} else if strings.EqualFold(returnType, Hours) {
		return diff.Hours(), nil
	} else if strings.EqualFold(returnType, Mins) {
		return diff.Minutes(), nil
	} else if strings.EqualFold(returnType, Seconds) {
		return diff.Seconds(), nil
	}
	return nil, fmt.Errorf("unknow return type [%s]", returnType)
}
