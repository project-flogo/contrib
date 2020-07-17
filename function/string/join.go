package string

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"strings"
)

func init() {
	_ = function.Register(&fnJoin{})
}

type fnJoin struct {
}

func (fnJoin) Name() string {
	return "join"
}

func (fnJoin) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeArray, data.TypeString}, false
}

func (fnJoin) Eval(params ...interface{}) (interface{}, error) {
	a, err := coerce.ToArray(params[0])
	if err != nil {
		return nil, fmt.Errorf("error converting string.join first argument [%+v] to array: %s", params[0], err.Error())
	}
	str, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("error converting string.join second argument [%+v] to string: %s", params[0], err.Error())
	}
	return strings.Join(toStringArray(a), str), nil
}

func toStringArray(array []interface{}) []string {
	var tmpArray []string
	for _, a := range array {
		str, _ := coerce.ToString(a)
		tmpArray = append(tmpArray, str)
	}
	return tmpArray
}
