package string

import (
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnReplaceAll{})
}

type fnReplaceAll struct {
}

func (fnReplaceAll) Name() string {
	return "replaceall"
}

func (fnReplaceAll) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString, data.TypeString}, false
}

func (fnReplaceAll) Eval(params ...interface{}) (interface{}, error) {
	return strings.ReplaceAll(params[0].(string), params[1].(string), params[2].(string)), nil
}
