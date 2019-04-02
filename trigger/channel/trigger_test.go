package channel

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/project-flogo/core/trigger"
	"github.com/stretchr/testify/assert"
)

var testMetadata *trigger.Metadata

//
func getTestMetadata(t *testing.T) *trigger.Metadata {

	if testMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("descriptor.json")
		assert.Nil(t, err)

		md := trigger.NewMetadata(string(jsonMetadataBytes))
		assert.NotNil(t, md)
		//
		testMetadata = md
	}

	return testMetadata
}

//
const testConfig string = `{
  "id": "flogo-channel",
  "ref": "github.com/project-flogo/contrib/trigger/channel",
  "handlers": [
    {
      "settings": {
        "channel": "test"
      },
      "action" : {
		"id": "dummy"
      }
    }
  ]
}`

func TestChannelFactory_New(t *testing.T) {

	f := &Factory{}

	config := &trigger.Config{}
	json.Unmarshal([]byte(testConfig), config)

	trg, err := f.New(config)
	assert.Nil(t, err)
	assert.NotNil(t, trg)
}

/*
func TestChannelTrigger_Initialize(t *testing.T) {
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

}


func TestChannelTrigger_Initialize(t *testing.T) {

	f := &Factory{}

	config := &trigger.Config{}
	err := json.Unmarshal([]byte(testConfig), config)
	assert.Nil(t, err)

	actions := map[string]action.Action{"dummy": test.NewDummyAction(func() {
		//do nothing
	})}

	trg, err := f.New(config)

	channels.Add("test:5")

	trg, err := trg.Initialize(f, config, actions)

	assert.Nil(t, err)
	assert.NotNil(t, trg)

	ct, ok := trg.(*Trigger)
	assert.True(t, ok)
	assert.Equal(t, 1, len(ct.handlers))

	channels.Close()
}
*/
//
//func TestChannelTrigger_Handler(t *testing.T) {
//	md := getTestMetadata(t)
//	f := &ChannelFactory{metadata: md}
//
//	config := &trigger.Config{}
//	err := json.Unmarshal([]byte(testConfig), config)
//	assert.Nil(t, err)
//
//	count := 0
//	actions := map[string]action.Action{"dummy": test.NewDummyAction(func() {
//		count++
//	})}
//
//	channels.Add("test:5")
//
//	trg, err := test.InitTrigger(f, config, actions)
//
//	assert.Nil(t, err)
//	assert.NotNil(t, trg)
//
//	ct, ok := trg.(*Trigger)
//	assert.True(t, ok)
//	assert.Equal(t, 1, len(ct.handlers))
//
//	err = trg.Start()
//	assert.Nil(t, err)
//
//	ch := channels.Get("test")
//	ch <- "val"
//
//	time.Sleep(100 * time.Millisecond)
//	assert.Equal(t, 1, count)
//
//	err = trg.Stop()
//	assert.Nil(t, err)
//
//	channels.Close()
//}
