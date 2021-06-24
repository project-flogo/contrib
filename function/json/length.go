package json

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnLength{})
}

type fnLength struct {
}

// Name returns the name of the function
func (fnLength) Name() string {
	return "length"
}

// Sig returns the function signature
func (fnLength) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny}, false
}

// Eval executes the function
func (fnLength) Eval(params ...interface{}) (interface{}, error) {
	switch t := params[0].(type) {
	case []interface{}:
		return len(t), nil
	}
	obj, err := coerce.ToObject(params[0])
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce [%+v] to object: %s", params[0], err.Error())
	}
	return len(obj), nil
}
