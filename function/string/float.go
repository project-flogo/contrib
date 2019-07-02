package string

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
	"math"
)

func init() {
	_ = function.Register(&fnFloat{})
}

type fnFloat struct {
}

// Name returns the name of the function
func (fnFloat) Name() string {
	return "float"
}

// Sig returns the function signature
func (fnFloat) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, true
}

// Eval fnFloat the function
func (s *fnFloat) Eval(ins ...interface{}) (interface{}, error) {

	log.RootLogger().Debugf("Start Float64 function with parameters %+v", ins)
	if len(ins) == 1 {
		return coerce.ToFloat64(ins[0])
	} else if len(ins) == 2 {
		val, err := coerce.ToFloat64(ins[0])
		if err != nil {
			return 0, fmt.Errorf("Invalid float type [%+v]", ins[0])
		}

		precision, err := coerce.ToInt(ins[1])
		if err != nil {
			return 0, fmt.Errorf("Float precision [%+v] must be integer", ins[1])
		}

		if precision > 16 {
			precision = 16
		}

		return getPrecisedFloat(val, precision), nil
	} else {
		return 0.0, fmt.Errorf("function arguments for float.float64 must be one or two")
	}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func getPrecisedFloat(num float64, precision int) float64 {
	output := math.Pow10(precision)
	return float64(round(num*output)) / output
}
