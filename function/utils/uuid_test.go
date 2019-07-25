package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var fnUuid = &fnUUID{}

func TestFnUUID_Eval(t *testing.T) {
	uuid, err := fnUuid.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, uuid)
}
