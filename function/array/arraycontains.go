package array

import (
	"github.com/project-flogo/core/support/log"
	"reflect"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

type Contains struct {
}

func init() {
	function.Register(&Contains{})
}

func (s *Contains) Name() string {
	return "contains"
}

func (Contains) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeAny}, false
}

func (s *Contains) Eval(params ...interface{}) (interface{}, error) {
	array := params[0]
	item := params[1]
	log.RootLogger().Infof("Looking for \"%s\" in \"%s\"", item, array)
	if array == nil || item == nil {
		return false, nil
	}
	arrV := reflect.ValueOf(array)
	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {
			if arrV.Index(i).Interface() == item {
				return true, nil
			}
		}
	}
	return false, nil
}
