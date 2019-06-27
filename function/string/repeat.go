package string

import (
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnRepeat{})
}

type fnRepeat struct {
}

func (fnRepeat) Name() string {
	return "repeat"
}

func (fnRepeat) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeInt}, false
}

func (fnRepeat) Eval(params ...interface{}) (interface{}, error) {

	return strings.Repeat(params[0].(string), params[1].(int)), nil
}
