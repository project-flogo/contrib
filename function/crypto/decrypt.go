package crypto

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnDecrypt{})
}

type fnDecrypt struct {
}

func (fnDecrypt) Name() string {
	return "decrypt"
}

func (fnDecrypt) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeBytes, data.TypeBytes}, false
}

func (fnDecrypt) Eval(params ...interface{}) (interface{}, error) {
	key, err := coerce.ToBytes(params[0])
	if err != nil {
		return nil, fmt.Errorf("decrypt function first parameter (key) [%+v] must be bytes", params[0])
	}

	ciphertext, err := coerce.ToBytes(params[1])
	if err != nil {
		return nil, fmt.Errorf("decrypt function second parameter (ciphertext) [%+v] must be byte", params[1])
	}

	return Decrypt(key, ciphertext)
}
