package string

import (
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnTrimRight{})
}

type fnTrimRight struct {
}

func (fnTrimRight) Name() string {
	return "trimRight"
}

func (fnTrimRight) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnTrimRight) Eval(params ...interface{}) (interface{}, error) {
	return strings.TrimRight(params[0].(string), params[1].(string)), nil
}
