package number

import (
	"github.com/project-flogo/core/data/coerce"
	"math/rand"
	"time"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnRandom{})
}

type fnRandom struct {
}

func (fnRandom) Name() string {
	return "random"
}

func (fnRandom) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeInt}, true
}

func (fnRandom) Eval(params ...interface{}) (interface{}, error) {
	limit := 10
	if len(params) > 0 {
		var err error
		limit, err = coerce.ToInt(params[0])
		if err != nil {
			limit = 10
		}
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(limit), nil
}
