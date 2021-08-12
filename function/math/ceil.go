package math

import (
	"fmt"
	"math"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnCeil{})
}

type fnCeil struct {
}

// Name returns the name of the function
func (fnCeil) Name() string {
	return "ceil"
}

// Sig returns the function signature
func (fnCeil) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeFloat64}, false
}

// Eval executes the function
func (fnCeil) Eval(params ...interface{}) (interface{}, error) {
	obj, err := coerce.ToFloat64(params[0])
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce [%+v] to float64: %s", params[0], err.Error())
	}
	return math.Ceil(obj), nil
}
