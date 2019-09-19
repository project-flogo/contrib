package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/project-flogo/core/data/coerce"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnDecodeBase64{})
}

type fnDecodeBase64 struct {
}

func (fnDecodeBase64) Name() string {
	return "decodeBase64"
}

func (fnDecodeBase64) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

// Eval - Decode base64 string
func (fnDecodeBase64) Eval(params ...interface{}) (interface{}, error) {
	s1, err := coerce.ToBytes(params[0])
	if err != nil {
		return nil, fmt.Errorf("decodeBase64 function first parameter [%+v] must be []byte", params[0])
	}

	dbuf := make([]byte, base64.StdEncoding.DecodedLen(len(s1)))
	n, err := base64.StdEncoding.Decode(dbuf, s1)
	return dbuf[:n], err
}
