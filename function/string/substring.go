package string

import (
	"fmt"
	"github.com/project-flogo/core/data/coerce"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnSubstring{})
}

type fnSubstring struct {
}

func (fnSubstring) Name() string {
	return "substring"
}

func (fnSubstring) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeInt, data.TypeInt}, false
}

func (fnSubstring) Eval(params ...interface{}) (interface{}, error) {
	str, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.substring function first parameter [%+v] must be string", params[0])
	}

	start, err := coerce.ToInt(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.substring function second parameter [%+v] must be integer", params[1])
	}

	length, err := coerce.ToInt(params[2])
	if err != nil {
		return nil, fmt.Errorf("string.substring function third parameter [%+v] must be integer", params[1])
	}

	if length == -1 {
		return str[start:], nil
	}

	if start+length > len(str) {
		return nil, fmt.Errorf("string length exceeded")
	}

	return str[start : start+length], nil
}
