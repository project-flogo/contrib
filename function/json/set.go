package json

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnSet{})
}

type fnSet struct {
}

// Name returns the name of the function
func (fnSet) Name() string {
	return "set"
}

// Sig returns the function signature
func (fnSet) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeObject, data.TypeString, data.TypeAny}, false
}

// Eval executes the function
func (fnSet) Eval(params ...interface{}) (interface{}, error) {
	switch params[0].(type) {
	case []interface{}:
		return nil, fmt.Errorf("Cannot set value for array type object using key")
	}
	obj, err := coerce.ToObject(params[0])
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce [%+v] to object: %s", params[0], err.Error())
	}
	key, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce the key to string: %s", err.Error())
	}
	val, err := coerce.ToAny(params[2])
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce the value to any: %s", err.Error())
	}
	if obj == nil {
		obj = make(map[string]interface{})
	}
	obj[key] = val
	return obj, nil
}
