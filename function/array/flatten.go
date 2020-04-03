package array

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"reflect"
)

type fnFlatten struct {
}

func init() {
	function.Register(&fnFlatten{})
}

func (a *fnFlatten) Name() string {
	return "flatten"
}

func (fnFlatten) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeArray, data.TypeInt}, false
}

func (fnFlatten) Eval(params ...interface{}) (interface{}, error) {
	if len(params) <= 0 {
		return nil, fmt.Errorf("array merge must have at least 1 array")
	}
	a, err := coerce.ToArray(params[0])
	if err != nil {
		return nil, fmt.Errorf("array.flatten's arg of dept must be array")
	}

	depth, err := coerce.ToInt(params[1])
	if err != nil {
		return nil, fmt.Errorf("array.flatten's arg of dept must be int")
	}

	if depth <= 0 {
		depth = -1
	}
	return flattern(a, 0, depth), nil
}

func isSlice(v interface{}) bool {
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Slice, reflect.Array:
		return true
	default:
		return false
	}
}

func flattern(args []interface{}, depth, max int) []interface{} {
	if max == 0 {
		//Do nothing
		return args
	}

	var list []interface{}
	depth++
	for _, v := range args {
		if isSlice(v) {
			if max == -1 || depth <= max {
				for _, z := range flattern((v.([]interface{})), depth, max) {
					list = append(list, z)
				}
			} else {
				list = append(list, v)
			}
		} else {
			list = append(list, v)
		}
	}
	return list

}
