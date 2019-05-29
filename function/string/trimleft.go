package string

import (
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnTrimLeft{})
}

type fnTrimLeft struct {
}

func (fnTrimLeft) Name() string {
	return "trimleft"
}

func (fnTrimLeft) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnTrimLeft) Eval(params ...interface{}) (interface{}, error) {
	return strings.TrimLeft(params[0].(string), params[1].(string)), nil
}
