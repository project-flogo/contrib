package string

import (
	"encoding/base64"
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/support/log"

	"github.com/project-flogo/core/data/expression/function"
)


type StringToBase64 struct {
}

func init() {
	function.Register(&StringToBase64{})
}

func (s *StringToBase64) Name() string {
	return "stringToBase64"
}

func (s *StringToBase64) GetCategory() string {
	return "string"
}

func (s *StringToBase64) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (s *StringToBase64) Eval(params ...interface{}) (interface{}, error) {

	str, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.stringToBase64 function first parameter [%+v] must be string", params[0])
	}
	log.RootLogger().Debugf("Returns base64 encoding of a given string \"%s\"", str)

	encoded := base64.StdEncoding.EncodeToString([]byte(str))

	return encoded, nil
}
