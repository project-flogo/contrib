package soapclient

import (
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestSettings(t *testing.T) {
	settings := &Settings{SoapServiceEndpoint: "https://www.dataaccess.com/webservicesserver/NumberConversion.wso", SoapVersion: "1.1"}
	iCtx := test.NewActivityInitContext(settings, nil)
	act, _ := New(iCtx)
	assert.NotNil(t, act)
}

func TestSOAPClientActivity(t *testing.T) {
	settings := &Settings{SoapServiceEndpoint: "https://www.dataaccess.com/webservicesserver/NumberConversion.wso", SoapVersion: "1.1", AttributePrefix: "@"}
	iCtx := test.NewActivityInitContext(settings, nil)
	act, _ := New(iCtx)
	tc := test.NewActivityContext(act.Metadata())

	//setup attrs
	input := &Input{SoapAction: "NumberToWords", RequestBody: map[string]interface{}{"NumberToWords": map[string]interface{}{"ubiNum": 100, "@xmlns": "http://www.dataaccess.com/webservicesserver/"}}}
	tc.SetInputObject(input)

	done, err := act.Eval(tc)
	assert.Nil(t, err)
	assert.True(t, done)

	output := &Output{}
	tc.GetOutputObject(output)
	assert.NotNil(t, output)
	assert.Equal(t, 200, output.HttpStatus)
	assert.False(t, output.IsFault)
	assert.NotNil(t, output.ResponsePayload)
	res, err1 := coerce.ToObject(output.ResponsePayload)
	assert.Nil(t, err1)
	resObj := res["NumberToWordsResponse"].(map[string]interface{})
	assert.Equal(t, "one hundred", resObj["NumberToWordsResult"])
}

func TestSOAPClientActivityFault(t *testing.T) {
	settings := &Settings{SoapServiceEndpoint: "http://www.dneonline.com/calculator.asmx", SoapVersion: "1.1", AttributePrefix: "@"}
	iCtx := test.NewActivityInitContext(settings, nil)
	act, _ := New(iCtx)
	tc := test.NewActivityContext(act.Metadata())

	//setup attrs
	input := &Input{RequestBody: map[string]interface{}{"Divide": map[string]interface{}{"intA": 6, "intB": 0}}}
	tc.SetInputObject(input)

	done, err := act.Eval(tc)
	assert.Nil(t, err)
	assert.True(t, done)

	output := &Output{}
	tc.GetOutputObject(output)
	assert.NotNil(t, output)
	assert.Equal(t, 500, output.HttpStatus)
	assert.True(t, output.IsFault)
	assert.NotNil(t, output.ResponseFault)
}