package string

import (
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnLastIndex{})
}

type fnLastIndex struct {
}

func (fnLastIndex) Name() string {
	return "lastIndex"
}

func (fnLastIndex) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnLastIndex) Eval(params ...interface{}) (interface{}, error) {
	return strings.LastIndex(params[0].(string), params[1].(string)), nil
}
