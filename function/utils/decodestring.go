package utils

import (
	"encoding/base64"
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnDecodeString{})
}

type fnDecodeString struct {
}

func (fnDecodeString) Name() string {
	return "decodestring"
}

func (fnDecodeString) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

// Eval - Decode base64 string
func (fnDecodeString) Eval(params ...interface{}) (interface{}, error) {
	data, err := base64.StdEncoding.DecodeString(params[0].(string))
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return string(data), nil
}
