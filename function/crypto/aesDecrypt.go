package crypto

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&aesDecryptFn{})
}

type aesDecryptFn struct {
}

// Name returns the name of the function
func (aesDecryptFn) Name() string {
	return "aesDecrypt"
}

// Sig returns the function signature
func (aesDecryptFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

// Eval executes the function
func (aesDecryptFn) Eval(params ...interface{}) (interface{}, error) {

	if logger.DebugEnabled() {
		logger.Debugf("Entering function aesEncrypt ()")
	}

	ciphertext, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("crypto.aesDecryptFn function first parameter [%+v] must be string", params[0])
	}

	key, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("crypto.aesDecryptFn function second parameter [%+v] must be string", params[1])
	}

	plaintext, err := aesDecrypt(ciphertext, []byte(key))
	if err != nil {
		return nil, err
	}

	if logger.DebugEnabled() {
		logger.Debugf("Exiting function aesEncrypt (eval)")
	}

	return string(plaintext), nil
}
