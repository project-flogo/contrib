package string

import (
	"fmt"
	"github.com/project-flogo/core/data/coerce"
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnRepeat{})
}

type fnRepeat struct {
}

func (fnRepeat) Name() string {
	return "repeat"
}

func (fnRepeat) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeInt}, false
}

func (fnRepeat) Eval(params ...interface{}) (interface{}, error) {
	s1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.repeat function first parameter [%+v] must be string", params[0])
	}
	s2, err := coerce.ToInt(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.repeat function second parameter [%+v] must be int", params[1])
	}
	return strings.Repeat(s1, s2), nil
}
