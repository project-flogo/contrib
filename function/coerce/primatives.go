package coerce

import (
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnToString{})
	function.Register(&fnToInt{})
	function.Register(&fnToInt32{})
	function.Register(&fnToInt64{})
	function.Register(&fnToFloat32{})
	function.Register(&fnToFloat64{})
	function.Register(&fnToBool{})
	function.Register(&fnToBytes{})
}

type fnToString struct {
	*baseFn
}

func (*fnToString) Name() string {
	return "toString"
}

func (*fnToString) Eval(params ...interface{}) (interface{}, error) {
	return coerce.ToString(params[0])
}

type fnToInt struct {
	*baseFn
}

func (*fnToInt) Name() string {
	return "toInt"
}

func (*fnToInt) Eval(params ...interface{}) (interface{}, error) {
	return coerce.ToInt(params[0])
}

type fnToInt32 struct {
	*baseFn
}

func (*fnToInt32) Name() string {
	return "toInt32"
}

func (*fnToInt32) Eval(params ...interface{}) (interface{}, error) {
	return coerce.ToInt32(params[0])
}

type fnToInt64 struct {
	*baseFn
}

func (*fnToInt64) Name() string {
	return "toInt64"
}

func (*fnToInt64) Eval(params ...interface{}) (interface{}, error) {
	return coerce.ToInt64(params[0])
}

type fnToFloat32 struct {
	*baseFn
}

func (*fnToFloat32) Name() string {
	return "toFloat32"
}

func (*fnToFloat32) Eval(params ...interface{}) (interface{}, error) {
	return coerce.ToFloat32(params[0])
}

type fnToFloat64 struct {
	*baseFn
}

func (*fnToFloat64) Name() string {
	return "toFloat64"
}

func (*fnToFloat64) Eval(params ...interface{}) (interface{}, error) {
	return coerce.ToFloat64(params[0])
}

type fnToBool struct {
	*baseFn
}

func (*fnToBool) Name() string {
	return "toBool"
}

func (*fnToBool) Eval(params ...interface{}) (interface{}, error) {
	return coerce.ToBool(params[0])
}

type fnToBytes struct {
	*baseFn
}

func (*fnToBytes) Name() string {
	return "toBytes"
}

func (*fnToBytes) Eval(params ...interface{}) (interface{}, error) {
	return coerce.ToBytes(params[0])
}
