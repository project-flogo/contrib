package json

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnValues{})
}

type fnValues struct {
}

// Name returns the name of the function
func (fnValues) Name() string {
	return "objValues"
}

// Sig returns the function signature
func (fnValues) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeObject}, false
}

// Eval executes the function
func (fnValues) Eval(params ...interface{}) (interface{}, error) {
	switch t := params[0].(type) {
	case []interface{}:
		return t, nil
	}
	input, err := coerce.ToObject(params[0])
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce [%+v] to object: %s", params[0], err.Error())
	}
	values := make([]interface{}, len(input))
	i := 0
	for _, v := range input {
		values[i] = v
		i++
	}
	return values, nil
}
