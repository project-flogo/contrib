package json

import (
	"encoding/json"
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

const testJsonData = `{
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
            },
            {
                "category": "fiction",
                "author": "Herman Melville",
                "title": "Moby Dick",
                "isbn": "0-553-21311-3",
                "price": 8.99
            },
            {
                "category": "fiction",
                "author": "J. R. R. Tolkien",
                "title": "The Lord of the Rings",
                "isbn": "0-395-19395-8",
                "price": 22.99
            }
        ],
        "bicycle": {
            "color": "red",
            "price": 19.95
        }
    },
    "expensive": 10
}`

func TestFnLen_Eval(t *testing.T) {
	var data interface{}
	err := json.Unmarshal([]byte(testJsonData), &data)
	assert.Nil(t, err)

	f := &fnPath{}
	v, err := function.Eval(f, "$.store.book[?(@.price == 22.99)].price[0]", data)
	assert.Nil(t, err)
	assert.Equal(t, 22.99, v)
}

func TestFnLoop_Eval(t *testing.T) {
	var data interface{}
	err := json.Unmarshal([]byte(testJsonData), &data)
	assert.Nil(t, err)

	f := &fnPath{}
	v, err := function.Eval(f, "$loop.store.book[?(@.price == 22.99)].price[0]", data)
	assert.Nil(t, err)
	assert.Equal(t, 22.99, v)
}
