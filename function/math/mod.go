package math

import (
	"fmt"
	"math"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnMod{})
}

type fnMod struct {
}

// Name returns the name of the function
func (fnMod) Name() string {
	return "mod"
}

// Sig returns the function signature
func (fnMod) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeFloat64, data.TypeFloat64}, false
}

// Eval executes the function
func (fnMod) Eval(params ...interface{}) (interface{}, error) {
	x, err := coerce.ToFloat64(params[0])
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce [%+v] to float64: %s", params[0], err.Error())
	}
	y, err := coerce.ToFloat64(params[1])
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce [%+v] to float64: %s", params[0], err.Error())
	}
	return math.Mod(x, y), nil
}
