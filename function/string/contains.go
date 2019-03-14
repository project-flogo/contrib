package string

import (
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

type fnContains struct {
}

func init() {
	function.Register(&fnContains{})
}

func (s *fnContains) Name() string {
	return "string.contains"
}

func (fnContains) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnContains) Eval(params ...interface{}) (interface{}, error) {
	str1 := params[0].(string)
	str2 := params[1].(string)
	return strings.Contains(str1, str2), nil
}
