package json

import (
	"encoding/json"
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

const inputExists = `{
    "store": {
        "book": [
            {
                "category": "reference",
                "author": "Nigel Rees",
                "title": "Sayings of the Century",
                "price": 8.95
            },
            {
                "category": "fiction",
                "author": "Evelyn Waugh",
                "title": "Sword of Honour",
                "price": 12.99
            }
        ],
        "bicycle": {
            "color": "red",
            "price": 19.95
        }
    },
    "expensive": 10,
	"emptyString": ""
}`

func TestFnCheckExists(t *testing.T) {
	var inputJSON interface{}
	err := json.Unmarshal([]byte(inputExists), &inputJSON)
	assert.Nil(t, err)

	f := &fnExists{}
	v, err := function.Eval(f, inputJSON, "$.store.book[?(@.price == 12.99)].price[0]")
	assert.Nil(t, err)
	assert.Equal(t, true, v)
}

func TestFnExistsLoop(t *testing.T) {
	var inputJSON interface{}
	err := json.Unmarshal([]byte(inputExists), &inputJSON)
	assert.Nil(t, err)

	f := &fnExists{}
	v, err := function.Eval(f, inputJSON, "$loop.store.book[?(@.price == 22.99)].price[0]")
	assert.Nil(t, err)
	assert.Equal(t, false, v)
}

func TestFnExistsNegative(t *testing.T) {
	var inputJSON interface{}
	err := json.Unmarshal([]byte(inputExists), &inputJSON)
	assert.Nil(t, err)

	f := &fnExists{}
	v, err := function.Eval(f, inputJSON, "$.store.abc")
	assert.Nil(t, err)
	assert.Equal(t, false, v)
}

func TestFnExistsEmpty(t *testing.T) {
	var inputJSON interface{}
	err := json.Unmarshal([]byte(inputExists), &inputJSON)
	assert.Nil(t, err)

	f := &fnExists{}
	v, err := function.Eval(f, inputJSON, "$.emptyString")
	assert.Nil(t, err)
	assert.Equal(t, true, v)
}

func TestFnExistsWithoutJSONPath(t *testing.T) {
	var inputJSON interface{}
	err := json.Unmarshal([]byte(inputExists), &inputJSON)
	assert.Nil(t, err)

	f := &fnExists{}
	v, err := function.Eval(f, inputJSON, "expensive")
	assert.Nil(t, err)
	assert.Equal(t, true, v)
}
