package string

import (
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnSplit{})
}

type fnSplit struct {
}

func (fnSplit) Name() string {
	return "split"
}

func (fnSplit) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnSplit) Eval(params ...interface{}) (interface{}, error) {
	return strings.Split(params[0].(string), params[1].(string)), nil
}
