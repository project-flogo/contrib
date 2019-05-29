package string

import (
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
	return "containsany"
}

func (fnContainsAny) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnContainsAny) Eval(params ...interface{}) (interface{}, error) {
	return strings.ContainsAny(params[0].(string), params[1].(string)), nil
}
