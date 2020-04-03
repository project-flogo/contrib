package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumFunc_Eval(t *testing.T) {

	tests := []struct {
		Array  []interface{}
		Result float64
	}{
		{
			Array:  []interface{}{1, 2, 3, 4, 5, 6, 7},
			Result: 28,
		},

		{
			Array:  []interface{}{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7},
			Result: 30.8,
		},
	}

	fn := &sumFunc{}
	for _, v := range tests {
		result, err := fn.Eval(v.Array)
		assert.Nil(t, err)
		assert.Equal(t, v.Result, result)
	}
}

func TestSliceFunc_Eval(t *testing.T) {

	tests := []struct {
		Array  []interface{}
		Start  int
		End    int
		Result []interface{}
	}{
		{
			Array:  []interface{}{1, 2, 3, 4, 5, 6, 7},
			Start:  1,
			End:    3,
			Result: []interface{}{2, 3},
		},
		{
			Array:  []interface{}{1, 2, 3, 4, 5, 6, 7},
			Start:  1,
			Result: []interface{}{2, 3, 4, 5, 6, 7},
		},
		{
			Array:  []interface{}{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7},
			Start:  0,
			End:    7,
			Result: []interface{}{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7},
		},
	}

	fn := &sliceFunc{}
	for _, v := range tests {
		result, err := fn.Eval(v.Array, v.Start, v.End)
		assert.Nil(t, err)
		assert.Equal(t, v.Result, result)
	}
}

func TestReverseFunc_Eval(t *testing.T) {

	tests := []struct {
		Array  []interface{}
		Result []interface{}
	}{
		{
			Array:  []interface{}{1, 2, 3, 4, 5, 6, 7},
			Result: []interface{}{7, 6, 5, 4, 3, 2, 1},
		},
		{
			Array:  []interface{}{5.5, 6.6, 7.7},
			Result: []interface{}{7.7, 6.6, 5.5},
		},
		{
			Array:  []interface{}{5.5, 6.6, 7.7, 8.8},
			Result: []interface{}{8.8, 7.7, 6.6, 5.5},
		},
	}

	fn := &reverseFunc{}
	for _, v := range tests {
		result, err := fn.Eval(v.Array)
		assert.Nil(t, err)
		assert.Equal(t, v.Result, result)
	}
}

func TestMergeFunc_Eval(t *testing.T) {
	fn := &mergeFunc{}

	obj := map[string]string{"key1": "value1", "key2": "value2"}
	obj2 := map[string]string{"key3": "value3", "key4": "value4"}

	tests := []struct {
		Array  []interface{}
		Array2 []interface{}
		Array3 []interface{}
		Result []interface{}
	}{
		{
			Array:  []interface{}{1, 2, 3, 4, 5, 6, 7},
			Array2: []interface{}{7, 6, 5, 4, 3, 2, 1},
			Result: []interface{}{1, 2, 3, 4, 5, 6, 7, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			Array:  []interface{}{5.5, 6.6, 7.7},
			Array2: []interface{}{7.7, 6.6, 5.5},
			Result: []interface{}{5.5, 6.6, 7.7, 7.7, 6.6, 5.5},
		},
		{
			Array:  []interface{}{obj, obj},
			Array2: []interface{}{obj2},
			Result: []interface{}{obj, obj, obj2},
		},
		{
			Array:  []interface{}{obj, obj},
			Array2: []interface{}{obj2},
			Array3: []interface{}{obj2},
			Result: []interface{}{obj, obj, obj2, obj2},
		},
	}

	for _, v := range tests {
		result, err := fn.Eval(v.Array, v.Array2, v.Array3)
		assert.Nil(t, err)
		assert.Equal(t, v.Result, result)
	}
}
