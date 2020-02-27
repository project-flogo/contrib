package array

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
)

type sumFunc struct {
}

func init() {
	function.Register(&sumFunc{})
}

func (a *sumFunc) Name() string {
	return "sum"
}

func (sumFunc) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny}, false
}

func (sumFunc) Eval(params ...interface{}) (interface{}, error) {
	arr := params[0]
	log.RootLogger().Debugf("Start array sum function with parameters %+v", arr)

	if arr == nil {
		//Do nothing
		return 0, nil
	}

	newArray, err := coerce.ToArray(arr)
	if err != nil {
		return nil, fmt.Errorf("array.sun function argument must be array")
	}

	sum := float64(0)

	for _, v := range newArray {
		num, err := coerce.ToFloat64(v)
		if err != nil {
			return nil, fmt.Errorf("array element must be number for array.sum function")
		}

		sum = sum + num
	}

	log.RootLogger().Debugf("array sum function done, final result %+v", sum)
	return sum, nil
}
