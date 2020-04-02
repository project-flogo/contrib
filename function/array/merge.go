package array

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
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

	if len(params) <= 0 {
		return nil, fmt.Errorf("array merge must have at least 1 array")
	}

	var interSlice []interface{}
	finalArrayValue := reflect.ValueOf(interSlice)
	if len(params) == 1 {
		//Do merge itself
		arrV := reflect.ValueOf(params[0])
		if arrV.Kind() == reflect.Slice {
			for i := 0; i < arrV.Len(); i++ {
				item := arrV.Index(i)
				if item.Kind() == reflect.Slice {
					for j := 0; j < item.Len(); j++ {
						finalArrayValue = reflect.Append(finalArrayValue, item.Index(j))
					}
				} else {
					a, err := coerce.ToArray(item.Interface())
					if err != nil {
						finalArrayValue = reflect.Append(finalArrayValue, arrV.Index(i))
					} else {
						for _, e := range a {
							finalArrayValue = reflect.Append(finalArrayValue, reflect.ValueOf(e))
						}
					}
				}
			}
		}
		return finalArrayValue.Interface(), nil
	}

	log.RootLogger().Debugf("Start array mergeFunc function with parameters %+v", params)
	for _, arg := range params {
		if arg != nil {
			arrV := reflect.ValueOf(arg)
			if arrV.Kind() == reflect.Slice {
				if !finalArrayValue.IsValid() {
					finalArrayValue = arrV
					continue
				} else {
					if arrV.Kind() == reflect.Slice {
						for i := 0; i < arrV.Len(); i++ {
							finalArrayValue = reflect.Append(finalArrayValue, arrV.Index(i))
						}
					} else {
						finalArrayValue = reflect.Append(finalArrayValue, arrV)
					}
				}

			}
		}
	}

	log.RootLogger().Debugf("array append function done, final array %+v", finalArrayValue.Interface())
	return finalArrayValue.Interface(), nil
}
