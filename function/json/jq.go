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

//var lock *sync.Mutex = &sync.Mutex{}

type jq struct {
}

// Name returns the name of the function
func (j jq) Name() string {
	return "jq"
}

// Sig returns the function signature
func (j jq) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeString}, false
}

// Eval executes the function
func (j jq) Eval(params ...interface{}) (interface{}, error) {
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

	input := copyInput(inputJSON)
	iter := query.Run(input)
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

// This function will deepcopy the array synchronously before passing to the jq evaluation.
// If the array entry is map then it will deepcopy the map in each array element.
// This is required because the same input is being modified internally in the jq code and simultaneously used in mapper
// code
func copyInput(input interface{}) interface{} {
	//lock.Lock()
	switch input := input.(type) {
	case []any:
		arr := make([]interface{}, len(input))
		for i, v := range input {
			switch v := v.(type) {
			case map[string]interface{}:
				copyMap := make(map[string]interface{})
				for rk, rv := range v {
					copyMap[rk] = rv
				}
				arr[i] = copyMap
			case interface{}:
				arr[i] = v
			}
		}
		//lock.Unlock()
		return arr
	}

	//lock.Unlock()
	return input
}
