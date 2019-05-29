package utils

import (
	"crypto/rand"
	"fmt"
	"io"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnUUID{})
}

type fnUUID struct {
}

func (fnUUID) Name() string {
	return "uuid"
}

func (fnUUID) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

// Eval - UUID generates a random UUID according to RFC 4122
func (fnUUID) Eval(params ...interface{}) (interface{}, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
