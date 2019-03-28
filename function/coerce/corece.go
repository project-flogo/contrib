package coerce

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnToType{})
}

type baseFn struct {
}

func (*baseFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny}, false
}

type fnToType struct {
}

func (*fnToType) Name() string {
	return "toType"
}

func (*fnToType) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeString}, false
}

func (*fnToType) Eval(params ...interface{}) (interface{}, error) {
	if len(params) < 2 {
		return nil, fmt.Errorf("missing params, signature is toType(any,string)")
	}

	typeStr, ok := params[1].(string)
	if !ok {
		return nil, fmt.Errorf("senond param must be a string")
	}

	t, err := data.ToTypeEnum(typeStr)
	if err != nil {
		return nil, err
	}

	return coerce.ToType(params[0], t)
}
