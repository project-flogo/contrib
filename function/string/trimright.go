package string

import (
	"fmt"
	"github.com/project-flogo/core/data/coerce"
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnTrimRight{})
}

type fnTrimRight struct {
}

func (fnTrimRight) Name() string {
	return "trimRight"
}

func (fnTrimRight) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnTrimRight) Eval(params ...interface{}) (interface{}, error) {

	s1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.trimRight function first parameter [%+v] must be string", params[0])
	}

	s2, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.trimRight function second parameter [%+v] must be string", params[1])
	}

	return strings.TrimRight(s1, s2), nil
}
