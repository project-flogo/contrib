package string

import (
	"fmt"
	"github.com/project-flogo/core/data/coerce"
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnReplace{})
}

type fnReplace struct {
}

func (fnReplace) Name() string {
	return "replace"
}

func (fnReplace) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString, data.TypeString, data.TypeInt}, false
}

func (fnReplace) Eval(params ...interface{}) (interface{}, error) {

	s1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.replace function first parameter [%+v] must be string", params[0])
	}
	s2, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.replace function second parameter [%+v] must be string", params[1])
	}
	s3, err := coerce.ToString(params[2])
	if err != nil {
		return nil, fmt.Errorf("string.replace function third parameter [%+v] must be string", params[2])
	}

	s4, err := coerce.ToInt(params[3])
	if err != nil {
		return nil, fmt.Errorf("string.replace function last parameter [%+v] must be int", params[3])
	}

	return strings.Replace(s1, s2, s3, s4), nil
}
