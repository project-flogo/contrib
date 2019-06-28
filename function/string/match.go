package string

import (
	"fmt"
	"github.com/project-flogo/core/data/coerce"
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
	s1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.matchRegEx function first parameter [%+v] must be string", params[0])
	}
	s2, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.matchRegEx function second parameter [%+v] must be string", params[1])
	}

	match, _ := regexp.MatchString(s1, s2)
	return match, nil
}
