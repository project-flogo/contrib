package string

import (
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnTrimPrefix{})
}

type fnTrimPrefix struct {
}

func (fnTrimPrefix) Name() string {
	return "trimPrefix"
}

func (fnTrimPrefix) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnTrimPrefix) Eval(params ...interface{}) (interface{}, error) {
	return strings.TrimPrefix(params[0].(string), params[1].(string)), nil
}
