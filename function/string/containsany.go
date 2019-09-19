package string

import (
	"fmt"
	"github.com/project-flogo/core/data/coerce"
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnContainsAny{})
}

type fnContainsAny struct {
}

func (fnContainsAny) Name() string {
	return "containsAny"
}

func (fnContainsAny) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnContainsAny) Eval(params ...interface{}) (interface{}, error) {

	str1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("containsAny function first parameter [%+v] must be string", params[0])
	}
	str2, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("containsAny function second parameter [%+v] must be string", params[1])
	}

	return strings.ContainsAny(str1, str2), nil
}
