package string

import (
	"fmt"
	"github.com/project-flogo/core/data/coerce"
	"strconv"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnInteger{})
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
	s1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.integer function first parameter [%+v] must be string", params[0])
	}
	return strconv.Atoi(s1)
}
