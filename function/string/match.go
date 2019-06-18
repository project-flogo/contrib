package string

import (
	"regexp"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnMatch{})
}

type fnMatch struct {
}

func (fnMatch) Name() string {
	return "matchRegEx"
}

func (fnMatch) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnMatch) Eval(params ...interface{}) (interface{}, error) {
	match, _ := regexp.MatchString(params[0].(string), params[1].(string))
	return match, nil
}
