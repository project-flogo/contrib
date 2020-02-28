package array

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
	"reflect"
)

type reverseFunc struct {
}

func init() {
	function.Register(&reverseFunc{})
}

func (a *reverseFunc) Name() string {
	return "reverse"
}

func (reverseFunc) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny}, false
}

func (reverseFunc) Eval(params ...interface{}) (interface{}, error) {
	items := params[0]
	log.RootLogger().Debugf("Start array reverseFunc function with parameters %+v", items)

	if items == nil {
		//Do nothing
		return items, nil
	}

	arrV := reflect.ValueOf(items)
	if arrV.Kind() == reflect.Slice {
		for i, j := 0, arrV.Len()-1; i < j; i, j = i+1, j-1 {
			left := arrV.Index(i).Interface()
			right := arrV.Index(j).Interface()
			arrV.Index(i).Set(reflect.ValueOf(right))
			arrV.Index(j).Set(reflect.ValueOf(left))
		}
	}

	log.RootLogger().Debugf("array append function done, final array %+v", arrV.Interface())
	return arrV.Interface(), nil
}
