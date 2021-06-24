package json

import (
	"strings"

	"github.com/oliveagle/jsonpath"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnPath{})
}

type fnPath struct {
}

// Name returns the name of the function
func (fnPath) Name() string {
	return "path"
}

// Sig returns the function signature
func (fnPath) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeAny}, false
}

// Eval executes the function
func (fnPath) Eval(params ...interface{}) (interface{}, error) {
	expression := params[0].(string)
	//tmp fix to take $loop as $. for now
	if strings.HasPrefix(strings.TrimSpace(expression), "$loop.") {
		expression = strings.Replace(expression, "$loop", "$", -1)
	}
	return jsonpath.JsonPathLookup(params[1], expression)
}
