package sample

/*
import (
	"encoding/json"
	"testing"

	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/support/test"
	"github.com/project-flogo/core/trigger"
	"github.com/stretchr/testify/assert"
)

const testConfig string = `{
	"id": "flogo-timer",
	"ref": "github.com/project-flogo/contrib/trigger/pulsar",
	"settings": {
	  "url": "pulsar://localhost:6650"
	},
	"handlers": [
	  {
			"action":{
				"id":"dummy"
			},
			"settings": {
			  "topic": "sample",
			  "subscription":"sample-1"
			}
	  }
	]

  }`

func TestPulsarTrigger_Initialize(t *testing.T) {
	f := &Factory{}

	config := &trigger.Config{}
	err := json.Unmarshal([]byte(testConfig), config)
	assert.Nil(t, err)

	actions := map[string]action.Action{"dummy": test.NewDummyAction(func() {
		//do nothing
	})}

	trg, err := test.InitTrigger(f, config, actions)
	assert.Nil(t, err)
	assert.NotNil(t, trg)

	err = trg.Start()
	assert.Nil(t, err)

}
*/
