package array

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"

	"reflect"
)

type Delete struct {
}

func init() {
	function.Register(&Delete{})
}

func (s *Delete) Name() string {
	return "delete"
}

func (s *Delete) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeInt}, false
}

func (s *Delete) Eval(params ...interface{}) (interface{}, error) {
	items := params[0]
	index, err := coerce.ToInt(params[1])
	if err != nil {
		return nil, fmt.Errorf("array delete function second parameter must be integer")
	}

	log.RootLogger().Debugf("Start array delete function with parameters %+v and %+v", items, index)

	if items == nil {
		return nil, fmt.Errorf("index out of bounds, index [%d] but array empty", index)
	}

	arrV := reflect.ValueOf(items)
	if arrV.Kind() != reflect.Slice {
		return nil, fmt.Errorf("unable to use array.delete on un-array object")
	}

	if index >= arrV.Len() {
		return nil, fmt.Errorf("index out of bounds, index [%d] but array length [%d]", index, arrV.Len())
	} else {
		before := arrV.Slice(0, index)
		after := arrV.Slice(index+1, arrV.Len())
		return reflect.AppendSlice(before, after).Interface(), nil
	}
}
