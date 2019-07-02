package string

import (
	"fmt"
	"github.com/project-flogo/core/data/coerce"
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnTrimPrefix{})
}

type fnTrimPrefix struct {
}

func (fnTrimPrefix) Name() string {
	return "trimPrefix"
}

func (fnTrimPrefix) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnTrimPrefix) Eval(params ...interface{}) (interface{}, error) {

	s1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.trimPrefix function first parameter [%+v] must be string", params[0])
	}

	s2, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.trimPrefix function second parameter [%+v] must be string", params[1])
	}

	return strings.TrimPrefix(s1, s2), nil
}
