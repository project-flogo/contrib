package string

import (
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnIndex{})
}

type fnIndex struct {
}

func (fnIndex) Name() string {
	return "index"
}

func (fnIndex) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnIndex) Eval(params ...interface{}) (interface{}, error) {
	return strings.Index(params[0].(string), params[1].(string)), nil
}
