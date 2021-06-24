package json

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnObjKeys{})
}

type fnObjKeys struct {
}

// Name returns the name of the function
func (fnObjKeys) Name() string {
	return "objKeys"
}

// Sig returns the function signature
func (fnObjKeys) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeObject}, false
}

// Eval executes the function
func (fnObjKeys) Eval(params ...interface{}) (interface{}, error) {
	switch params[0].(type) {
	case []interface{}:
		return nil, fmt.Errorf("Cannot list keys for array type object")
	}
	input, err := coerce.ToObject(params[0])
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce [%+v] to object: %s", params[0], err.Error())
	}
	keys := make([]string, len(input))
	i := 0
	for k := range input {
		keys[i] = k
		i++
	}
	return keys, nil
}
