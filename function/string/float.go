package string

import (
	"strconv"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnFloat{})
}

type fnFloat struct {
}

// Name returns the name of the function
func (fnFloat) Name() string {
	return "string.float"
}

// Sig returns the function signature
func (fnFloat) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

// Eval fnFloat the function
func (fnFloat) Eval(params ...interface{}) (interface{}, error) {
	return strconv.ParseFloat(params[0].(string), 64)
}
