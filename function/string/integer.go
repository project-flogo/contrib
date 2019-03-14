package string

import (
	"strconv"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnInteger{})
}

type fnInteger struct {
}

// Name returns the name of the function
func (fnInteger) Name() string {
	return "integer"
}

// Sig returns the function signature
func (fnInteger) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

// Eval fnInteger the function
func (fnInteger) Eval(params ...interface{}) (interface{}, error) {
	return strconv.Atoi(params[0].(string))
}
