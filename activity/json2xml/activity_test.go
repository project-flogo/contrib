package json2xml

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/clbanning/mxj"
	"github.com/project-flogo/core/data/coerce"
	"math"
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestStructure(t *testing.T) {

	type NestedData struct {
		DataA string `json:"data-a" xml:"data-a"`
	}

	type Data struct {
		Name       string     `json:"name" xml:"name"`
		Special    string     `json:"spec_chars" xml:"spec_chars"`
		List       []string   `json:"list" xml:"list"`
		Number     int        `json:"number" xml:"number"`
		Decimal    float32    `json:"decimal" xml:"decimal"`
		Utf8String string     `json:"utf8" xml:"utf8"`
		Nested     NestedData `json:"nested" xml:"nested"`
	}

	var rawTestData = &Data{
		Name:       "Test Name",
		Special:    "<>&\"'",
		List:       []string{"one", "two", "three"},
		Number:     43,
		Decimal:    math.Pi,
		Utf8String: "नमस्ते दुनिया", // Hello World in Hindi
		Nested: NestedData{
			DataA: "nested data value",
		},
	}

	ac := &Activity{}
	tc := test.NewActivityContext(ac.Metadata())

	// given
	var jsonObject, _ = coerce.ToObject(rawTestData) // json object type for Flogo framework
	t.Logf("Marshaled JSON: %s", jsonObject)

	acInput := &Input{Json: jsonObject, XmlRootTag: "Data"}
	tc.SetInputObject(acInput)

	// when
	done, _ := ac.Eval(tc)

	// then
	acOutput := &Output{}
	tc.GetOutputObject(acOutput)
	t.Logf("Marshaled XML: %s", string(acOutput.Xml))

	var outputData = &Data{}
	var err = xml.Unmarshal(acOutput.Xml, &outputData)
	if err != nil {
		t.Error(err)
		return
	}

	assert.True(t, done)
	assert.Equal(t, rawTestData, outputData)
}

func TestEvalNoXMLRoot(t *testing.T) {
	// setup
	ac := &Activity{}
	tc := test.NewActivityContext(ac.Metadata())

	// given
	jsonString := `{"hello": "world"}`
	var jsonObject map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonObject)

	acInput := &Input{Json: jsonObject, XmlRootTag: ""}
	tc.SetInputObject(acInput)

	// when
	done, _ := ac.Eval(tc)

	// then
	acOutput := &Output{}
	tc.GetOutputObject(acOutput)

	assert.True(t, done)
	var expectedResult = []byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<hello>world</hello>")
	assert.Equal(t, expectedResult, acOutput.Xml)
}

func TestEvalCustomXMLRoot(t *testing.T) {
	// setup
	ac := &Activity{}
	tc := test.NewActivityContext(ac.Metadata())

	// given
	jsonString := `{"hello": "world"}`
	var jsonObject map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonObject)

	acInput := &Input{Json: jsonObject, XmlRootTag: "my"}
	tc.SetInputObject(acInput)

	// when
	done, _ := ac.Eval(tc)

	// then
	acOutput := &Output{}
	tc.GetOutputObject(acOutput)

	assert.True(t, done)
	assert.Contains(t, string(acOutput.Xml), "my")
}

func TestJsonContainingXML(t *testing.T) {
	// setup
	ac := &Activity{}
	tc := test.NewActivityContext(ac.Metadata())

	// given
	jsonString := `{"xml": "<hello>world</hello>"}`
	var jsonObject map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonObject)

	acInput := &Input{Json: jsonObject}
	tc.SetInputObject(acInput)

	// when
	done, _ := ac.Eval(tc)

	// then
	expectedResult :=
		`<?xml version="1.0" encoding="UTF-8"?>` +
			"\n" +
			`<xml>&lt;hello&gt;world&lt;/hello&gt;</xml>`
	acOutput := &Output{}
	tc.GetOutputObject(acOutput)

	assert.True(t, done)
	assert.Equal(t, expectedResult, string(acOutput.Xml))
}

func TestEmptyCustomXMLRootTag(t *testing.T) {
	// setup
	ac := &Activity{}
	tc := test.NewActivityContext(ac.Metadata())

	// given
	acInput := &Input{Json: nil, XmlRootTag: "root"}
	tc.SetInputObject(acInput)

	// when
	done, _ := ac.Eval(tc)

	// then
	expectedResult :=
		`<?xml version="1.0" encoding="UTF-8"?>` +
			"\n" +
			`<root/>`
	acOutput := &Output{}
	tc.GetOutputObject(acOutput)

	assert.True(t, done)
	assert.Equal(t, expectedResult, string(acOutput.Xml))
}

func TestEmptyDefaultXMLRootTag(t *testing.T) {
	// setup
	ac := &Activity{}
	tc := test.NewActivityContext(ac.Metadata())

	// given
	acInput := &Input{Json: nil}
	tc.SetInputObject(acInput)

	// when
	done, _ := ac.Eval(tc)

	// then
	expectedResult :=
		`<?xml version="1.0" encoding="UTF-8"?>` +
			"\n" +
			fmt.Sprintf("<%s/>", mxj.DefaultRootTag)
	acOutput := &Output{}
	tc.GetOutputObject(acOutput)

	assert.True(t, done)
	assert.Equal(t, expectedResult, string(acOutput.Xml))
}
