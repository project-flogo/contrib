package activity_mapper

import (
	"encoding/json"
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestSimpleMapper(t *testing.T) {

	//set mappings
	mappingsJson :=
		`{
           "Output1": "1",
           "Output2": 2
		}`

	var mappings interface{}
	err := json.Unmarshal([]byte(mappingsJson), &mappings)
	if err != nil {
		panic("Unable to parse mappings: " + err.Error())
	}

	settings := map[string]interface{}{"mappings": mappings}

	act, err := New(settings)
	assert.Nil(t, err)

	ah := newActivityHost()
	tc := test.NewActivityContextWithAction(act.Metadata(), ah)

	//eval
	act.Eval(tc)

	//assert.Nil(t, ah.ReplyErr)
	o1, exists1 := ah.HostData.GetValue("Output1")
	assert.True(t, exists1, "Output1 not set")
	if exists1 {
		assert.Equal(t, "1", o1)
	}
	o2, exists2 := ah.HostData.GetValue("Output2")
	assert.True(t, exists2, "Output2 not set")
	if exists2 {
		assert.Equal(t, 2, o2)
	}
}

func newActivityHost() *test.TestActivityHost {
	input := map[string]data.TypedValue{"Input1": data.NewTypedValue(data.TypeString, "")}
	output := map[string]data.TypedValue{"Output1": data.NewTypedValue(data.TypeString, ""), "Output2": data.NewTypedValue(data.TypeInt, "")}

	ac := &test.TestActivityHost{
		HostId:     "1",
		HostRef:    "github.com/TIBCOSoftware/flogo-contrib/action/flow",
		IoMetadata: &metadata.IOMetadata{Input: input, Output: output},
		HostData:   data.NewSimpleScope(nil, nil),
	}

	return ac
}
