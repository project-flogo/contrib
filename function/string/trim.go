package string

import (
	"fmt"
	"github.com/project-flogo/core/data/coerce"
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnTrim{})
}

type fnTrim struct {
}

func (fnTrim) Name() string {
	return "trim"
}

func (fnTrim) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnTrim) Eval(params ...interface{}) (interface{}, error) {
	s1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.trim function first parameter [%+v] must be string", params[0])
	}

	s2, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.trim function second parameter [%+v] must be string", params[1])
	}
	return strings.Trim(s1, s2), nil
}
