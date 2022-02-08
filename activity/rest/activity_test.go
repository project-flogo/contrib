package rest

import (
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/resolve"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestSettings(t *testing.T) {
	settings := &Settings{UseEnvProp: "NO", Method: "POS", Uri: "http://petstore.swagger.io/v2/pet"}

	iCtx := test.NewActivityInitContext(settings, nil)
	_, err := New(iCtx)
	assert.NotNil(t, err)

	settings = &Settings{UseEnvProp: "YES", Method: "POS", Uri: ""}

	iCtx = test.NewActivityInitContext(settings, nil)
	_, err = New(iCtx)
	assert.NotNil(t, err)

	settings = &Settings{UseEnvProp: "NO", Method: "pOsT", Uri: "http://petstore.swagger.io/v2/pet"}

	iCtx = test.NewActivityInitContext(settings, nil)
	_, err = New(iCtx)
	assert.Nil(t, err)

	settings = &Settings{UseEnvProp: "YES", Method: "pOsT", Uri: "http://petstore.swagger.io/v2/pet"}

	iCtx = test.NewActivityInitContext(settings, nil)
	_, err = New(iCtx)
	assert.Nil(t, err)
}

const reqPostStr string = `{
  "name": "my pet"
}
`

var petID string

func TestSimplePost(t *testing.T) {

	settings := &Settings{UseEnvProp: "NO", Method: "POST", Uri: "http://petstore.swagger.io/v2/pet"}

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(settings, mf)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	//setup attrs
	tc.SetInput("content", reqPostStr)

	//eval
	act.Eval(tc)
	assert.NotNil(t, tc.GetOutput("data"))

}

func TestPassThrough(t *testing.T) {

	settings := &Settings{UseEnvProp: "NO", Method: "TRIGGER", Uri: "http://petstore.swagger.io/v2/pet"}

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(settings, mf)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	//setup attrs
	tc.SetInput("content", reqPostStr)
	tc.SetInput("method", "GET")

	//eval
	act.Eval(tc)
	assert.NotNil(t, tc.GetOutput("data"))

}

func TestPassThroughEnv(t *testing.T) {

	// url to hit: http://petstore.swagger.io/v2/pet

	settings := &Settings{UseEnvProp: "YES", Method: "TRIGGER", Uri: "{envProp}/:restOfPath"}

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(settings, mf)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	//setup attrs
	pathParams := map[string]string{
		"restOfPath": "v2/pet",
	}
	tc.SetInput("content", reqPostStr)
	tc.SetInput("method", "GET")
	tc.SetInput("envPropUri", "http://petstore.swagger.io")
	tc.SetInput("pathParams", pathParams)

	//eval
	act.Eval(tc)
	assert.NotNil(t, tc.GetOutput("data"))

}

func TestSimpleGet(t *testing.T) {

	settings := &Settings{UseEnvProp: "YES", Method: "GET", Uri: "{prop}/v2/pet/16"}

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(settings, mf)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("envPropUri", "http://www.httpbin.org/anything/firstPart")
	//eval
	act.Eval(tc)

	assert.NotNil(t, tc.GetOutput("data"))

}

func TestSimpleGetWithHeaders(t *testing.T) {

	settings := &Settings{UseEnvProp: "YES", Method: "GET", Uri: "http://petstore.swagger.io/v2/pet/1"}

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(settings, mf)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	headers := make(map[string]string)
	headers["TestHeader"] = "TestValue"
	tc.SetInput("headers", headers)
	tc.SetInput("envPropUri", "http://www.httpbin.org/anything/firstPart")
	//eval
	act.Eval(tc)

	output := &Output{}
	tc.GetOutputObject(output)
	assert.NotNil(t, output.Status)

}

func TestParamGet(t *testing.T) {

	settings := &Settings{UseEnvProp: "YES", Method: "GET", Uri: "http://petstore.swagger.io/v2/pet/:id"}

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(settings, mf)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	pathParams := map[string]string{
		"id": petID,
	}
	tc.SetInput("pathParams", pathParams)

	//eval
	act.Eval(tc)

	assert.NotNil(t, tc.GetOutput("data"))

}

func TestSimpleGetQP(t *testing.T) {

	settings := &Settings{UseEnvProp: "YES", Method: "GET", Uri: "http://petstore.swagger.io/v2/pet/findByStatus"}

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(settings, mf)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	queryParams := map[string]string{
		"status": "ava",
	}
	tc.SetInput("queryParams", queryParams)

	//eval
	act.Eval(tc)

	assert.NotNil(t, tc.GetOutput("data"))
}

func TestBuildURI(t *testing.T) {

	uri := "http://localhost:7070/flow/:id"

	params := map[string]string{
		"id": "1234",
	}

	newURI, _ := BuildURI(uri, params)

	assert.NotNil(t, newURI)
}

func TestBuildURI2(t *testing.T) {

	uri := "https://127.0.0.1:7070/:cmd/:id/test"

	params := map[string]string{
		"cmd": "flow",
		"id":  "1234",
	}

	newURI, _ := BuildURI(uri, params)
	assert.NotNil(t, newURI)
}

func TestBuildURI3(t *testing.T) {

	uri := "http://:localhost/flow/:id"

	params := map[string]string{
		"id":        "1234",
		"localhost": "google.com",
	}

	newURI, _ := BuildURI(uri, params)

	assert.NotNil(t, newURI)
}

func TestBuildURI4(t *testing.T) {

	uri := "https://127.0.0.1/:cmd/:id/test"

	params := map[string]string{
		"cmd": "flow",
		"id":  "1234",
	}

	newURI, _ := BuildURI(uri, params)

	assert.NotNil(t, newURI)
}

func TestReplaceSubStringNothingToReplace(t *testing.T) {

	uri := "{startdomainname}/:cmd/:id/test"

	result := replaceSubString(uri, "http://www.google.com")
	assert.Contains(t, result, "http://www.google.com/:cmd/:id/test")
}

func TestReplaceSubString(t *testing.T) {

	uri := "{startdomainname}/:cmd/:id/test"

	result := replaceSubString(uri, "http://www.google.com")
	assert.Contains(t, result, "http://www.google.com/:cmd/:id/test")
}

func TestMe(t *testing.T) {
	e := "X-Autorisation"
	a := e[2:]
	assert.Equal(t, "Autorisation", a)
}
