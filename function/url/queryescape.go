package url

import (
	"fmt"
	"net/url"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnQueryEscape{})
}

type fnQueryEscape struct {
}

// Name returns the name of the function
func (fnQueryEscape) Name() string {
	return "queryEscape"
}

// Sig returns the function signature
func (fnQueryEscape) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

// Eval executes the function
func (fnQueryEscape) Eval(params ...interface{}) (interface{}, error) {
	rawString, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce [%+v] to string: %s", params[0], err.Error())
	}
	return url.QueryEscape(rawString), nil
}
