package array

import (
	"fmt"
	"github.com/project-flogo/core/data/resolve"
	"reflect"
	"testing"

	"github.com/project-flogo/core/data/expression/script"
	"github.com/stretchr/testify/assert"
)

var resolver = resolve.NewCompositeResolver(map[string]resolve.Resolver{"static": nil, ".": nil, "env": &resolve.EnvResolver{}})
var factory = script.NewExprFactory(resolver)

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

func TestArrayEmtpy(t *testing.T) {
	final, _ := s.Eval(nil, "Mouse")
	for _, item := range final.([]string) {
		assert.Equal(t, item, "Mouse")
	}
}
