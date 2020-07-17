package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var c = &fnRewriteCookies{}

func TestRewriteCookies(t *testing.T) {
	cookies := []interface{}{"lo=; Path=/; Expires=Wed, 01 Jan 1800 00:00:00 GMT; Secure", "tsc=UcRKVG68jC...; Domain=cloud.tibco.com; Path=/; Max-Age=86399; HttpOnly; Secure"}
	final, err := c.Eval(cookies, "tsc", "mashery.com", "/mashery-path/")
	assert.Nil(t, err)
	assert.NotNil(t, final)
	expectedResult := []interface{}{"lo=; Path=/; Expires=Wed, 01 Jan 1800 00:00:00 GMT; Secure", "tsc=UcRKVG68jC...; Domain=mashery.com; Path=/mashery-path/; Max-Age=86399; HttpOnly; Secure"}
	assert.Equal(t, expectedResult, final)
}

func TestReqCookieToParams(t *testing.T) {
	cookies := "lo=; Path=/; Expires=Wed, 01 Jan 1800 00:00:00 GMT; Secure"
	f := &fnReqCookiesToParams{}
	final, err := f.Eval(cookies)
	assert.Nil(t, err)
	assert.Equal(t, "/", final.(map[string]string)["Path"])
	assert.NotNil(t, final)
}

func TestReqCookieToParamsFromParam(t *testing.T) {
	cookies := "lo=; Path=/; Expires=Wed, 01 Jan 1800 00:00:00 GMT; Secure"
	f := &fnReqCookiesToParams{}
	final, err := f.Eval(cookies)
	assert.Nil(t, err)
	assert.Equal(t, "/", final.(map[string]string)["Path"])

	f2 := &fnReqCookiesFromParams{}
	value, err := f2.Eval(final)
	assert.Nil(t, err)
	assert.NotNil(t, value)
}

func TestReqCookieToObject(t *testing.T) {
	cookies := "lo=; Path=/; Expires=Wed, 01 Jan 1800 00:00:00 GMT; Secure"
	f := &fnResCookieToObject{}
	final, err := f.Eval(cookies)
	assert.Nil(t, err)
	assert.NotNil(t, final)

	f2 := &fnResCookieFromObject{}
	value, err := f2.Eval(final)
	assert.Nil(t, err)
	assert.NotNil(t, value)
}

func TestResCookiesToObjectMap(t *testing.T) {
	cookies := "lo=; Path=/; Expires=Wed, 01 Jan 1800 00:00:00 GMT; Secure"
	f := &fnResCookiesToObjectMap{}
	final, err := f.Eval(cookies)
	assert.Nil(t, err)
	assert.NotNil(t, final)

	f2 := &fnResCookiesFromObjectMap{}
	value, err := f2.Eval(final)
	assert.Nil(t, err)
	assert.NotNil(t, value)
}
