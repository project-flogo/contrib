package string

import (
	"regexp"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnReplaceregex{})
}

type fnReplaceregex struct {
}

func (fnReplaceregex) Name() string {
	return "replaceregex"
}

func (fnReplaceregex) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString, data.TypeString}, false
}

func (fnReplaceregex) Eval(params ...interface{}) (interface{}, error) {
	re := regexp.MustCompile(params[0].(string))
	return re.ReplaceAllString(params[1].(string), params[2].(string)), nil
}
