package array

import (
	"fmt"
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
	return []data.Type{data.TypeAny, data.TypeAny}, true
}

func (mergeFunc) Eval(params ...interface{}) (interface{}, error) {
	if len(params) < 2 {
		return nil, fmt.Errorf("array merge must have at least 2 arrays")
	}

	log.RootLogger().Debugf("Start array mergeFunc function with parameters %+v and %+v", params)

	finalArrayValue := reflect.Value{}
	for _, arg := range params {
		if arg != nil {
			arrV := reflect.ValueOf(arg)
			if arrV.Kind() == reflect.Slice {
				if !finalArrayValue.IsValid() {
					finalArrayValue = arrV
					continue
				} else {
					item := reflect.ValueOf(arg)
					if item.Kind() == reflect.Slice {
						for i := 0; i < item.Len(); i++ {
							finalArrayValue = reflect.Append(finalArrayValue, item.Index(i))
						}
					} else {
						finalArrayValue = reflect.Append(finalArrayValue, item)
					}
				}

			}
		}
	}

	log.RootLogger().Debugf("array append function done, final array %+v", finalArrayValue.Interface())
	return finalArrayValue.Interface(), nil
}
