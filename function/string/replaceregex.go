package string

import (
	"fmt"
	"github.com/project-flogo/core/data/coerce"
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
	return "replaceRegEx"
}

func (fnReplaceregex) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString, data.TypeString}, false
}

func (fnReplaceregex) Eval(params ...interface{}) (interface{}, error) {

	s1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.replaceRegEx function first parameter [%+v] must be string", params[0])
	}
	s2, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.replaceRegEx function second parameter [%+v] must be string", params[1])
	}
	s3, err := coerce.ToString(params[2])
	if err != nil {
		return nil, fmt.Errorf("string.replaceRegEx function third parameter [%+v] must be string", params[2])
	}
	re := regexp.MustCompile(s1)
	return re.ReplaceAllString(s2, s3), nil
}
