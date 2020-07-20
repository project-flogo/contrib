package string

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"strings"
)

func init() {
	_ = function.Register(&fnTitle{})
}

type fnTitle struct {
}

func (fnTitle) Name() string {
	return "toTitleCase"
}

func (fnTitle) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (fnTitle) Eval(params ...interface{}) (interface{}, error) {
	str, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("error converting string.toTitleCase's argument [%+v] to string: %s", params[0], err.Error())
	}
	return strings.Title(str), nil
}
