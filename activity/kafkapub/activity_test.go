package kafkapub

import (
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&KafkaPubActivity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}
