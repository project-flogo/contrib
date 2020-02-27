package array

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
	"reflect"
)

type mergeFunc struct {
}

func init() {
	function.Register(&mergeFunc{})
}

func (a *mergeFunc) Name() string {
	return "merge"
}

func (mergeFunc) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeAny}, false
}

func (mergeFunc) Eval(params ...interface{}) (interface{}, error) {
	items := params[0]
	items2 := params[1]

	log.RootLogger().Debugf("Start array mergeFunc function with parameters %+v and %+v", items, items2)

	if items == nil {
		//Do nothing
		return items2, nil
	} else if items2 == nil {
		return items, nil
	}

	arrV := reflect.ValueOf(items)
	if arrV.Kind() == reflect.Slice {
		item := reflect.ValueOf(items2)
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
