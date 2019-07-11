package array

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var s = &appendFunc{}

func TestStaticAppend(t *testing.T) {
	Original := []string{"Cat", "Dog", "Snake"}
	expectedResult := []string{"Cat", "Dog", "Snake", "Mouse"}
	final, _ := s.Eval(Original, "Mouse")
	fmt.Println(reflect.TypeOf(final))
	fmt.Println(final)
	for i, item := range final.([]string) {
		assert.Equal(t, item, expectedResult[i])
	}
}

func TestAppendArray(t *testing.T) {
	Original := []string{"Cat", "Dog", "Snake"}
	target := []string{"Cat", "Dog", "Snake"}
	expectedResult := []string{"Cat", "Dog", "Snake", "Cat", "Dog", "Snake"}
	final, _ := s.Eval(Original, target)
	for i, item := range final.([]string) {
		assert.Equal(t, item, expectedResult[i])
	}
}

func TestArrayEmtpy(t *testing.T) {
	final, _ := s.Eval(nil, "Mouse")
	for _, item := range final.([]string) {
		assert.Equal(t, item, "Mouse")
	}
}
