package timer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitOk(t *testing.T) {
	f := &Factory{}
	tgr, err := f.New(nil)
	assert.Nil(t, err)
	assert.NotNil(t, tgr)
}
