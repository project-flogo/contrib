package math

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnIsNaN{})
}

type fnIsNaN struct {
}

// Name returns the name of the function
func (fnIsNaN) Name() string {
	return "isNaN"
}

// Sig returns the function signature
func (fnIsNaN) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeFloat64}, false
}

// Eval executes the function
func (fnIsNaN) Eval(params ...interface{}) (interface{}, error) {
	input, err := coerceToFloat64(params[0])
	if err != nil {
		return true, err
	}
	return math.IsNaN(input), nil
}

func coerceToFloat64(input interface{}) (float64, error) {
	switch t := input.(type) {
	case float32:
		return float64(t), nil
	case float64:
		return t, nil
	case int:
		return float64(t), nil
	case int8:
		return float64(t), nil
	case int32:
		return float64(t), nil
	case int64:
		return float64(t), nil
	case uint:
		return float64(t), nil
	case uint8:
		return float64(t), nil
	case uint16:
		return float64(t), nil
	case uint32:
		return float64(t), nil
	case uint64:
		return float64(t), nil
	case json.Number:
		return t.Float64()
	default:
		return 0.0, fmt.Errorf("Unable to coerce %#v to float64", input)
	}
}
