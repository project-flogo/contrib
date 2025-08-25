package datetime

import (
	"time"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&Timestamp{})
}

type Timestamp struct {
}

// Name returns the name of the function
func (Timestamp) Name() string {
	return "timestamp"
}

// Sig returns the function signature
func (Timestamp) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

// Eval executes the function
func (Timestamp) Eval(params ...interface{}) (interface{}, error) {

	outputTimeInMillis := time.Now().UnixNano() / int64(time.Millisecond)

	return outputTimeInMillis, nil
}
