package http

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnReqCookiesToParams{})
	_ = function.Register(&fnReqCookiesFromParams{})
	_ = function.Register(&fnResCookieToObject{})
	_ = function.Register(&fnResCookieFromObject{})
	_ = function.Register(&fnResCookiesToObjectMap{})
	_ = function.Register(&fnResCookiesFromObjectMap{})
}

type fnReqCookiesToParams struct {
}

func (fnReqCookiesToParams) Name() string {
	return "reqCookieToParams"
}

func (fnReqCookiesToParams) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (fnReqCookiesToParams) Eval(params ...interface{}) (interface{}, error) {

	cookies, err := coerce.ToString(params[0])
	if err != nil {
		return nil, err
	}

	//dummy request to parse cookie string
	r := &http.Request{}
	r.Header.Set("Cookie", cookies)
	cos := r.Cookies()
	cAsParams := make(map[string]string, len(cos))
	for _, cookie := range cos {
		cAsParams[cookie.Name] = cookie.Value
	}

	return cAsParams, nil
}

type fnReqCookiesFromParams struct {
}

func (fnReqCookiesFromParams) Name() string {
	return "reqCookieFromParams"
}

func (fnReqCookiesFromParams) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeParams}, false
}

func (fnReqCookiesFromParams) Eval(params ...interface{}) (interface{}, error) {

	cAsParams, err := coerce.ToParams(params[0])
	if err != nil {
		return nil, err
	}

	//dummy request to construct cookie string
	r := &http.Request{}

	for name, value := range cAsParams {
		r.AddCookie(&http.Cookie{Name:name, Value:value})
	}

	cookies := r.Header.Get("Cookie")

	return cookies, nil
}


type fnResCookieToObject struct {

}

func (fnResCookieToObject) Name() string {
	return "resCookieToObject"
}

func (fnResCookieToObject) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (fnResCookieToObject) Eval(params ...interface{}) (interface{}, error) {

	cookie, err := coerce.ToString(params[0])
	if err != nil {
		return nil, err
	}

	//dummy request to parse cookie string
	r := &http.Response{}
	r.Header.Set("Set-Cookie", cookie)
	cos := r.Cookies()

	if len(cos) > 0 {
		return cos[0], nil
	}

	return nil, nil
}

type fnResCookieFromObject struct {
}

func (fnResCookieFromObject) Name() string {
	return "resCookieFromObject"
}

func (fnResCookieFromObject) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny}, false
}

func (fnResCookieFromObject) Eval(params ...interface{}) (interface{}, error) {

	if len(params) == 0 {
		return nil, nil
	}

	return toCookieString(params[0])
}

type fnResCookiesToObjectMap struct {
}

func (fnResCookiesToObjectMap) Name() string {
	return "resCookiesToObjectMap"
}

func (fnResCookiesToObjectMap) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeArray}, false
}

func (fnResCookiesToObjectMap) Eval(params ...interface{}) (interface{}, error) {

	if len(params) == 0 {
		return nil, nil
	}

	cookies, err := coerce.ToArray(params[0])
	if err != nil {
		return nil, err
	}

	//dummy response to parse cookie string
	r := &http.Response{}
	for _, cookie := range cookies {

		cookieStr, err := coerce.ToString(cookie)
		if err != nil {
			return nil, err
		}
		r.Header.Add("Set-Cookie", cookieStr)
	}

	cos := r.Cookies()
	coMap := make(map[string]interface{}, len(cos))

	for _, c := range cos {
		coMap[c.Name] = c
	}

	return coMap, nil
}

type fnResCookiesFromObjectMap struct {
}

func (fnResCookiesFromObjectMap) Name() string {
	return "resCookiesFromObjectMap"
}

func (fnResCookiesFromObjectMap) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeMap}, false
}

func (fnResCookiesFromObjectMap) Eval(params ...interface{}) (interface{}, error) {

	if len(params) == 0 {
		return nil, nil
	}

	coMap, err := coerce.ToObject(params[0])
	if err != nil {
		return nil, err
	}

	var cookies []interface{}

	for _, co := range coMap {
		cookie, err := toCookieString(co)
		if err != nil {
			return nil, err
		}

		cookies = append(cookies, cookie)
	}

	return cookies, nil
}

func toCookieString(co interface{}) (string, error){
	strCookie := ""

	if c, ok := co.(*http.Cookie); ok {
		strCookie = c.String()
	} else if c, ok := co.(map[string]interface{}); ok {
		cookie, err := mapToCookie(c)
		if err != nil {
			return "", err
		}
		strCookie = cookie.String()
	} else if params, ok := co.(map[string]string); ok {
		c, _ := coerce.ToObject(params)

		cookie, err := mapToCookie(c)
		if err != nil {
			return "", err
		}
		strCookie = cookie.String()
	} else {
		return "", fmt.Errorf("unsupported cookie format: %v", co)
	}

	return strCookie, nil
}


func mapToCookie(values map[string]interface{}) (cookie *http.Cookie, err error) {

	cookie = &http.Cookie{}

	for key, value := range values {
		switch strings.ToLower(key) {
		case "name":
			cookie.Name = value.(string)
		case "value":
			cookie.Value, err = coerce.ToString(value)
			if err != nil {
				return nil, err
			}
		case "path":
			cookie.Path = value.(string)
		case "domain":
			cookie.Domain = value.(string)
		case "expires":
			if dt, ok := value.(time.Time); ok {
				cookie.Expires = dt.UTC()
			} else if dt, ok := value.(string); ok {
				exptime, err := time.Parse(time.RFC1123, dt)
				if err != nil {
					exptime, err = time.Parse("Mon, 02-Jan-2006 15:04:05 MST", dt)
					if err != nil {
						return nil, fmt.Errorf("invalid expiration: %s", dt)
					}
				}
				cookie.Expires = exptime.UTC()
			}
		case "maxage":
			cookie.MaxAge, err = coerce.ToInt(value)
			if err != nil {
				return nil, err
			}
		case "secure":
			cookie.Secure, err = coerce.ToBool(value)
			if err != nil {
				return nil, err
			}
		case "httponly":
			cookie.HttpOnly, err = coerce.ToBool(value)
			if err != nil {
				return nil, err
			}
		case "samesite":
			if ss, err := coerce.ToInt(value); err == nil {
				cookie.SameSite = http.SameSite(ss)
			} else if ss, ok := value.(string); ok {
				switch strings.ToLower(ss) {
				case "lax":
					cookie.SameSite = http.SameSiteLaxMode
				case "strict":
					cookie.SameSite = http.SameSiteStrictMode
				default:
					cookie.SameSite = http.SameSiteDefaultMode
				}
			}
		}

	}

	return cookie, nil
}
