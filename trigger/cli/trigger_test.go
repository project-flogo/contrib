package cli

import (
	"encoding/json"
	"testing"

	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/support/test"
	"github.com/project-flogo/core/trigger"
	"github.com/stretchr/testify/assert"
)

//var jsonTestMetadata = getTestJsonMetadata()
//
//func getTestJsonMetadata() string {
//	jsonMetadataBytes, err := ioutil.ReadFile("trigger.json")
//	if err != nil {
//		panic("No Json Metadata found for trigger.json path")
//	}
//	return string(jsonMetadataBytes)
//}

const testConfig string = `{
  "id": "flogo-cli",
	"ref": "github.com/project-flogo/contrib/trigger/cli",
	"settings":{
		"singleCmd":true
	},
  "handlers": [
    {
		"action":{
			"id": "dummy"
		},
		"settings": {
        "command": "run"
      }
    }
  ]
}
`

func TestInitOk(t *testing.T) {
	// New  factory
	f := &Factory{}

	config := &trigger.Config{}
	json.Unmarshal([]byte(testConfig), &config)
	tgr, err := f.New(config)
	assert.Nil(t, err)
	assert.NotNil(t, tgr)

}

func TestCliTrigger_Initialize(t *testing.T) {
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
	_, err = Invoke()
	assert.Nil(t, err)
}

/*
//TODO fix this test
func TestInitOk(t *testing.T) {
	// New  factory
	f := &CliTriggerFactory{}
	tgr := f.New("flogo-cli")

	runner := &TestRunner{}

	config := trigger.Config{}
	json.Unmarshal([]byte(testConfig), &config)
	tgr.Init(config, runner)
}
*/

/*
//TODO fix this test
func TestHandlerOk(t *testing.T) {

	// New  factory
	f := &CliTriggerFactory{}
	tgr := f.New("flogo-cli")

	runner := &TestRunner{}

	config := trigger.Config{}
	tgr.Init(runner)

	tgr.Start()
	defer tgr.Stop()

	uri := "http://127.0.0.1:8091/device/12345/reset"

	req, err := http.NewRequest("POST", uri, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Debug("response Status:", resp.Status)

	if resp.StatusCode >= 300 {
		t.Fail()
	}
}
*/
