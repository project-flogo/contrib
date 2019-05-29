package utils

import (
	"encoding/base64"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnEncodeString{})
}

type fnEncodeString struct {
}

func (fnEncodeString) Name() string {
	return "encodestring"
}

func (fnEncodeString) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

// Eval - UUID generates a random UUID according to RFC 4122
func (fnEncodeString) Eval(params ...interface{}) (interface{}, error) {
	data := []byte(params[0].(string))
	str := base64.StdEncoding.EncodeToString(data)
	return str, nil
}
