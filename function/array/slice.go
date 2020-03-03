package array

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
	"reflect"
)

type sliceFunc struct {
}

func init() {
	function.Register(&sliceFunc{})
}

func (a *sliceFunc) Name() string {
	return "slice"
}

func (sliceFunc) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeInt}, true
}

func (sliceFunc) Eval(params ...interface{}) (interface{}, error) {
	array := params[0]
	log.RootLogger().Debugf("Start array slice function with parameters %+v", array)
	if array == nil {
		//Do nothing
		return array, nil
	}

	start, err := coerce.ToInt(params[1])
	if err != nil {
		return nil, fmt.Errorf("array slice second arguments must be a integer")
	}

	arrV := reflect.ValueOf(array)
	if arrV.Kind() == reflect.Slice {
		var le int
		if len(params) >= 3 {
			le, err = coerce.ToInt(params[2])
			if err != nil {
				return nil, fmt.Errorf("array slice third arguments must be a integer")
			}
			if le == 0 {
				le = arrV.Len()
			}
		} else {
			le = arrV.Len()
		}

		if arrV.Len() < le {
			return nil, fmt.Errorf("array slice end index out of bound")
		}

		v := arrV.Slice(start, le)
		log.RootLogger().Debugf("array slice function done, final array %+v", v.Interface())
		return v.Interface(), nil
	}
	return nil, fmt.Errorf("array slice first argument must be array")

}
