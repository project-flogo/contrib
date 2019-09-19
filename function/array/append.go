package array

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
	"reflect"
)

type appendFunc struct {
}

func init() {
	function.Register(&appendFunc{})
}

func (a *appendFunc) Name() string {
	return "append"
}

func (appendFunc) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeAny}, false
}

func (appendFunc) Eval(params ...interface{}) (interface{}, error) {
	items := params[0]
	item := params[1]
	log.RootLogger().Debugf("Start array appendFunc function with parameters %+v and %+v", items, item)

	if item == nil {
		//Do nothing
		return items, nil
	}

	if items == nil {
		newitems := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(item)), 1, 1)
		newitems.Index(0).Set(reflect.ValueOf(item))
		return newitems.Interface(), nil
	}

	arrV := reflect.ValueOf(items)
	if arrV.Kind() == reflect.Slice {
		item := reflect.ValueOf(item)
		if item.Kind() == reflect.Slice {
			for i := 0; i < item.Len(); i++ {
				arrV = reflect.Append(arrV, item.Index(i))
			}
		} else {
			arrV = reflect.Append(arrV, item)
		}
	}

	log.RootLogger().Debugf("array append function done, final array %+v", arrV.Interface())
	return arrV.Interface(), nil
}
