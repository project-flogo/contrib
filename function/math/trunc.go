package math

import (
	"fmt"
	"math"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnTrunc{})
}

type fnTrunc struct {
}

// Name returns the name of the function
func (fnTrunc) Name() string {
	return "trunc"
}

// Sig returns the function signature
func (fnTrunc) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeFloat64}, false
}

// Eval executes the function
func (fnTrunc) Eval(params ...interface{}) (interface{}, error) {
	obj, err := coerce.ToFloat64(params[0])
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce [%+v] to float64: %s", params[0], err.Error())
	}
	return math.Trunc(obj), nil
}
