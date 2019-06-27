package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/project-flogo/core/data/coerce"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnEncodeBase64{})
}

type fnEncodeBase64 struct {
}

func (fnEncodeBase64) Name() string {
	return "encodeBase64"
}

func (fnEncodeBase64) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeBytes}, false
}

// Eval - UUID generates a random UUID according to RFC 4122
func (fnEncodeBase64) Eval(params ...interface{}) (interface{}, error) {
	s1, err := coerce.ToBytes(params[0])
	if err != nil {
		return nil, fmt.Errorf("encodeBase64 function first parameter [%+v] must be []byte", params[0])
	}

	ebuf := make([]byte, base64.StdEncoding.EncodedLen(len(s1)))
	base64.StdEncoding.Encode(ebuf, s1)
	return ebuf, err
}
