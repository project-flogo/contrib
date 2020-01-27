package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var d = &Delete{}

func TestStaticDelete(t *testing.T) {
	expectedResult := []string{"Cat", "Dog", "Snake"}
	final, err := d.Eval(expectedResult, 2)
	assert.Nil(t, err)
	assert.Equal(t, []string{"Cat", "Dog"}, final)
}
