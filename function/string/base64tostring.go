package string

import (
	"encoding/base64"
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/support/log"

	"github.com/project-flogo/core/data/expression/function"
)

type Base64ToString struct {
}

func init() {
	function.Register(&Base64ToString{})
}

func (s *Base64ToString) Name() string {
	return "base64ToString"
}

func (s *Base64ToString) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (s *Base64ToString) Eval(params ...interface{}) (interface{}, error) {

	base64str, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("Base64ToString function argument must be string")
	}
	decoded, err := base64.StdEncoding.DecodeString(base64str)
	if err != nil {
		log.RootLogger().Debugf("decode error:", err)
		return "", err
	}
	return string(decoded), err
}
