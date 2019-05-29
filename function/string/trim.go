package string

import (
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
	return strings.Trim(params[0].(string), params[1].(string)), nil
}
