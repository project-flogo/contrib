package coerce

import (
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnToParams{})
	_ = function.Register(&fnToObject{})
	_ = function.Register(&fnToArray{})
}

type fnToParams struct {
	*baseFn
}

func (*fnToParams) Name() string {
	return "toParams"
}

func (*fnToParams) Eval(params ...interface{}) (interface{}, error) {
	return coerce.ToParams(params[0])
}

type fnToObject struct {
	*baseFn
}

func (*fnToObject) Name() string {
	return "toObject"
}

func (*fnToObject) Eval(params ...interface{}) (interface{}, error) {
	return coerce.ToObject(params[0])
}

type fnToArray struct {
	*baseFn
}

func (*fnToArray) Name() string {
	return "toArray"
}

func (*fnToArray) Eval(params ...interface{}) (interface{}, error) {
	return coerce.ToArray(params[0])
}
