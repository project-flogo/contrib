package url

import (
	"fmt"
	"net/url"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnQuery{})
}

type fnQuery struct {
}

// Name returns the name of the function
func (fnQuery) Name() string {
	return "query"
}

// Sig returns the function signature
func (fnQuery) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeBool}, false
}

// Eval executes the function
func (fnQuery) Eval(params ...interface{}) (interface{}, error) {
	rawString, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce [%+v] to string: %s", params[0], err.Error())
	}
	u, err := url.Parse(rawString)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse url: %s", err.Error())
	}
	encode, err := coerce.ToBool(params[1])
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce [%+v] to bool: %s", params[1], err.Error())
	}
	if encode {
		return u.Query().Encode(), nil
	}
	return u.Query(), nil
}
