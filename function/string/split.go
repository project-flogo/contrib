package string

import (
	"fmt"
	"github.com/project-flogo/core/data/coerce"
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

	s1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.split function first parameter [%+v] must be string", params[0])
	}
	s2, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.split function second parameter [%+v] must be string", params[1])
	}
	return strings.Split(s1, s2), nil
}
