package url

import (
	"fmt"
	"net/url"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnEscapePath{})
}

type fnEscapePath struct {
}

// Name returns the name of the function
func (fnEscapePath) Name() string {
	return "escapePath"
}

// Sig returns the function signature
func (fnEscapePath) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

// Eval executes the function
func (fnEscapePath) Eval(params ...interface{}) (interface{}, error) {
	rawString, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce [%+v] to string: %s", params[0], err.Error())
	}
	u, err := url.Parse(rawString)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse url: %s", err.Error())
	}
	return u.EscapedPath(), nil
}
