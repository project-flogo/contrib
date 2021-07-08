package json

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnNumbersToString{})
}

type fnNumbersToString struct {
}

// Name returns the name of the function
func (fnNumbersToString) Name() string {
	return "numbersToString"
}

// Sig returns the function signature
func (fnNumbersToString) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny}, false
}

// Eval executes the function
func (fnNumbersToString) Eval(params ...interface{}) (interface{}, error) {
	inputBytes, err := json.Marshal(params[0])
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(inputBytes)
	decoder := json.NewDecoder(reader)
	decoder.UseNumber()
	switch t := params[0].(type) {
	case []interface{}:
		outputArr := make([]interface{}, len(t))
		err = decoder.Decode(&outputArr)
		if err != nil {
			return nil, err
		}
		return handleArray(outputArr), nil
	case map[string]interface{}:
		outputMap := make(map[string]interface{})
		err = decoder.Decode(&outputMap)
		if err != nil {
			return nil, err
		}
		encodeNumbersToString(outputMap)
		return outputMap, nil
	default:
		return nil, fmt.Errorf("Unsupported json object type [%T]", params[0])
	}
}

func encodeNumbersToString(m map[string]interface{}) {
	for k, v := range m {
		switch t := v.(type) {
		case json.Number:
			m[k] = t.String()
		case map[string]interface{}:
			encodeNumbersToString(t)
		case []interface{}:
			m[k] = handleArray(t)
		default:
			fmt.Printf("Unsupported type: %T\n", v)
		}
	}
}

func handleArray(arr []interface{}) []interface{} {
	for i, v := range arr {
		switch t := v.(type) {
		case json.Number:
			arr[i] = t.String()
		case map[string]interface{}:
			encodeNumbersToString(t)
		case []interface{}:
			arr[i] = handleArray(t)
		default:
			fmt.Printf("Unsupported type inside array: %T\n", v)
		}
	}
	return arr
}
