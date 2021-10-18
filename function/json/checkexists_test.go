package json

import (
	"encoding/json"
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

const inputCheckExists = `{
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
    "expensive": 10
}`

func TestFnCheckExists(t *testing.T) {
	var inputJSON interface{}
	err := json.Unmarshal([]byte(inputCheckExists), &inputJSON)
	assert.Nil(t, err)

	f := &fnCheckExists{}
	v, err := function.Eval(f, "$.store.book[?(@.price == 22.95)].price[0]", inputJSON)
	assert.Nil(t, err)
	assert.Equal(t, true, v)
}

func TestFnCheckExistsLoop(t *testing.T) {
	var inputJSON interface{}
	err := json.Unmarshal([]byte(inputCheckExists), &inputJSON)
	assert.Nil(t, err)

	f := &fnCheckExists{}
	v, err := function.Eval(f, "$loop.store.book[?(@.price == 22.99)].price[0]", inputJSON)
	assert.Nil(t, err)
	assert.Equal(t, false, v)
}

func TestFnCheckExistsNegative(t *testing.T) {
	var inputJSON interface{}
	err := json.Unmarshal([]byte(inputCheckExists), &inputJSON)
	assert.Nil(t, err)

	f := &fnCheckExists{}
	v, err := function.Eval(f, "$.store.abc", inputJSON)
	assert.Nil(t, err)
	assert.Equal(t, false, v)
}
