package array

import (
	"encoding/json"
	"fmt"
	"github.com/project-flogo/core/data/coerce"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatternFunc(t *testing.T) {

	fn := &fnFlatten{}

	str := `[
	 [
	   {
	     "id": 1
	   }
	 ],
	 [
	   {
	     "id": 2
	   },
	   {
	     "id": 3
	   }
	 ]
	]`

	var d interface{}
	err := json.Unmarshal([]byte(str), &d)
	assert.Nil(t, err)
	final, err := fn.Eval(d, -1)
	assert.Nil(t, err)
	print(final)
	assert.Equal(t, 3, len(final.([]interface{})))

	obj := map[string]string{"key1": "value1", "key2": "value2"}

	var aa = []interface{}{[]interface{}{obj}, []interface{}{obj}, []interface{}{obj}}
	final, err = fn.Eval(aa, -1)
	assert.Nil(t, err)
	print(final)
	assert.Equal(t, 3, len(final.([]interface{})))

	obj = map[string]string{"key1": "value1", "key2": "value2"}
	aa = []interface{}{obj}
	final, err = fn.Eval(aa, -1)
	assert.Nil(t, err)
	print(final)
	assert.Equal(t, aa, final)

	str = "[1, 2, [3, 4, [5, 6, [7, 8, [9, 10]]]]]"
	a, _ := coerce.ToArray(str)
	final, err = fn.Eval(a, -1)
	assert.Nil(t, err)
	print(final)
	assert.Equal(t, 10, len(final.([]interface{})))

	str = "[1, 2, [3, 4, [5, 6, [7, 8, [9, 10]]]]]"
	a, _ = coerce.ToArray(str)
	final, err = fn.Eval(a, 2)
	assert.Nil(t, err)
	print(final)
	assert.Equal(t, 7, len(final.([]interface{})))
}

func print(in interface{}) {
	v, _ := json.Marshal(in)
	fmt.Println(string(v))
}
