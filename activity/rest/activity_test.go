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
	settings := &Settings{Method: "POS", Uri: "http://petstore.swagger.io/v2/pet"}

	iCtx := test.NewActivityInitContext(settings, nil)
	_, err := New(iCtx)
	assert.NotNil(t, err)

	settings = &Settings{Method: "POST", Uri: ""}

	iCtx = test.NewActivityInitContext(settings, nil)
	_, err = New(iCtx)
	assert.NotNil(t, err)

	settings = &Settings{Method: "pOsT", Uri: "http://petstore.swagger.io/v2/pet"}

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

	settings := &Settings{Method: "POST", Uri: "http://petstore.swagger.io/v2/pet"}

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

func TestSimpleGet(t *testing.T) {

	settings := &Settings{Method: "GET", Uri: "http://petstore.swagger.io/v2/pet/16"}

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(settings, mf)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	//eval
	act.Eval(tc)

	assert.NotNil(t, tc.GetOutput("data"))

}

func TestSimpleGetWithHeaders(t *testing.T) {

	settings := &Settings{Method: "GET", Uri: "http://petstore.swagger.io/v2/pet/1"}

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(settings, mf)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	headers := make(map[string]string)
	headers["TestHeader"] = "TestValue"
	tc.SetInput("headers", headers)

	//eval
	act.Eval(tc)

	output := &Output{}
	tc.GetOutputObject(output)
	assert.NotNil(t, output.Status)

}

func TestParamGet(t *testing.T) {

	settings := &Settings{Method: "GET", Uri: "http://petstore.swagger.io/v2/pet/:id"}

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

//func TestSimpleGetWithProxy(t *testing.T) {
//
//	settings := &Settings{Method:"GET", Uri:"http://petstore.swagger.io/v2/pet/1"}
//	settings.Proxy = "http://localhost:12345"
//
//	mf := mapper.NewFactory(resolve.GetBasicResolver())
//	iCtx := test.NewActivityInitContext(settings, mf)
//	act, err := New(iCtx)
//	assert.Nil(t, err)
//
//	tc := test.NewActivityContext(act.Metadata())
//
//	//eval
//	_, err = act.Eval(tc)
//	if err != nil {
//		fmt.Printf("error: %v\n", err)
//	}
//	val := tc.GetOutput("result")
//	fmt.Printf("result: %v\n", val)
//}

func TestSimpleGetQP(t *testing.T) {

	settings := &Settings{Method: "GET", Uri: "http://petstore.swagger.io/v2/pet/findByStatus"}

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

	newURI := BuildURI(uri, params)

	assert.NotNil(t, newURI)
}

func TestBuildURI2(t *testing.T) {

	uri := "https://127.0.0.1:7070/:cmd/:id/test"

	params := map[string]string{
		"cmd": "flow",
		"id":  "1234",
	}

	newURI := BuildURI(uri, params)
	assert.NotNil(t, newURI)
}

func TestBuildURI3(t *testing.T) {

	uri := "http://localhost/flow/:id"

	params := map[string]string{
		"id": "1234",
	}

	newURI := BuildURI(uri, params)

	assert.NotNil(t, newURI)
}

func TestBuildURI4(t *testing.T) {

	uri := "https://127.0.0.1/:cmd/:id/test"

	params := map[string]string{
		"cmd": "flow",
		"id":  "1234",
	}

	newURI := BuildURI(uri, params)

	assert.NotNil(t, newURI)
}
