package json

import (
	"fmt"
	"strings"

	"github.com/oliveagle/jsonpath"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnExists{})
}

type fnExists struct {
}

// Name returns the name of the function
func (fnExists) Name() string {
	return "exists"
}

// Sig returns the function signature
func (fnExists) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeString}, false
}

// Eval executes the function
func (fnExists) Eval(params ...interface{}) (interface{}, error) {
	expression, ok := params[1].(string)
	if !ok {
		return false, fmt.Errorf("The JSON key/path must be a string")
	}
	//tmp fix to take $loop as $. for now
	if strings.HasPrefix(strings.TrimSpace(expression), "$loop.") {
		expression = strings.Replace(expression, "$loop", "$", -1)
	}
	if !strings.HasPrefix(expression, "$.") {
		expression = "$." + expression
	}
	_, err := jsonpath.JsonPathLookup(params[0], expression)
	if err != nil {
		return false, nil
	}
	return true, nil
}
