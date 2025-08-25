package crypto

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&hmacFn{})
}

type hmacFn struct {
}

// Name returns the name of the function
func (hmacFn) Name() string {
	return "hmac"
}

// Sig returns the function signature
func (hmacFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString,data.TypeString}, false
}

// Eval executes the function
func (hmacFn) Eval(params ...interface{}) (interface{}, error) {

	if logger.DebugEnabled() {
		logger.Debugf("Entering function hmac(eval)")
	}

	value, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("crypto.hmacFn function first parameter [%+v] must be string", params[0])
	}

	key, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("crypto.hmacFn function second parameter [%+v] must be string", params[1])
	}


	hmac := hmacValue([]byte(value), []byte(key))


	if logger.DebugEnabled() {
		logger.Debugf("Exiting function hmac(eval)")
	}

	return hmac, nil
}
