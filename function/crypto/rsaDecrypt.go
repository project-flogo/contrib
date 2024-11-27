package crypto

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&rsaDecryptFn{})
}

type rsaDecryptFn struct {
}

// Name returns the name of the function
func (rsaDecryptFn) Name() string {
	return "rsaDecrypt"
}

// Sig returns the function signature
func (rsaDecryptFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

// Eval executes the function
func (rsaDecryptFn) Eval(params ...interface{}) (interface{}, error) {

	if logger.DebugEnabled() {
		logger.Debugf("Entering function rsaDecrypt ()")
	}

	ciphertextAsBase64, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("crypto.rsaDecrypt function first parameter [%+v] must be string", params[0])
	}

	key, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("crypto.rsaDecrypt function second parameter [%+v] must be string", params[1])
	}

	ciphertext, err := decodeBase64(ciphertextAsBase64)
	if err!= nil {
        return nil, err
    }

	plaintext, err := rsaDecrypt(ciphertext, []byte(key))
	if err != nil {
		return nil, err
	}

	if logger.DebugEnabled() {
		logger.Debugf("Exiting function rsaDecrypt()")
	}

	return string(plaintext), nil
}
