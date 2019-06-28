package string

import (
	"fmt"
	"github.com/project-flogo/core/data/coerce"
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnContains{})
}

type fnContains struct {
}

func (s *fnContains) Name() string {
	return "contains"
}

func (fnContains) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnContains) Eval(params ...interface{}) (interface{}, error) {
	str, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.contains function first parameter [%+v] must be string", params[0])
	}
	substr, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.contains function second parameter [%+v] must be string", params[1])
	}
	return strings.Contains(str, substr), nil

}
