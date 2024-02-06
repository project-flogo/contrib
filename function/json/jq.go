package json

import (
	"fmt"

	"github.com/itchyny/gojq"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&jq{})
}

type jq struct {
}

// Name returns the name of the function
func (jq) Name() string {
	return "jq"
}

// Sig returns the function signature
func (jq) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeString}, false
}

// Eval executes the function
func (jq) Eval(params ...interface{}) (interface{}, error) {

	var err error
	var inputJSON interface{}

	//1st parameter will be the input json
	inputJSON, err = coerce.ToAny(params[0])
	if err != nil {
		return nil, fmt.Errorf("unable to coerce input JSON [%+v]:  %s", params[0], err.Error())

	}

	//2nd parameter will be the query
	inputQuery, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("unable to coerce the query to string: %s", err.Error())
	}
	query, err := gojq.Parse(inputQuery)
	if err != nil {
		return nil, err
	}
	var result []interface{}
	iter := query.Run(inputJSON)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return nil, fmt.Errorf("error running the query : %s", err)
		}
		result = append(result, v)
	}

	return result, nil
}
