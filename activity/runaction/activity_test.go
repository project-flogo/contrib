package runaction

import (
	"github.com/project-flogo/core/activity"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}
