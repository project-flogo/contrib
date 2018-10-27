package rest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/logger"
)

var log = logger.GetLogger("activity-rest")

func init() {
	activity.Register(&Activity{}, New)
}

const (
	methodPOST  = "POST"
	methodPUT   = "PUT"
	methodPATCH = "PATCH"
)

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	act := &Activity{settings: s}
	act.containsParam = strings.Index(s.Uri, "/:") > -1

	httpTransportSettings := http.Transport{}

	// Set the proxy server to use, if supplied
	if len(s.Proxy) > 0 {
		proxyURL, err := url.Parse(s.Proxy)
		if err != nil {
			log.Debugf("Error parsing proxy url '%s': %s", s.Proxy, err)
			return nil, err
		}

		if log.DebugEnabled() {
			log.Debug("Setting proxy server:", s.Proxy)
		}
		httpTransportSettings.Proxy = http.ProxyURL(proxyURL)
	}

	// Skip ssl validation
	if s.SkipSSL {
		httpTransportSettings.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	act.client = http.Client{Transport: &httpTransportSettings}

	return act, nil
}

// Activity is an activity that is used to invoke a REST Operation
// settings : {method, uri, headers, proxy, skipSSL}
// input    : {pathParams, queryParams, headers, content}
// outputs  : {status, result}
type Activity struct {
	settings      *Settings
	containsParam bool
	client        http.Client
}

func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Invokes a REST Operation
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	ctx.GetInputObject(input)

	uri := a.settings.Uri

	if a.containsParam {

		if len(input.PathParams) == 0 {
			err := activity.NewError("Path Params not specified, required for URI: "+uri, "", nil)
			return false, err
		}

		uri = BuildURI(a.settings.Uri, input.PathParams)
	}

	if len(input.QueryParams) > 0 {
		qp := url.Values{}

		for key, value := range input.QueryParams {
			qp.Set(key, value)
		}

		uri = uri + "?" + qp.Encode()
	}

	if log.DebugEnabled() {
		log.Debugf("REST Call: [%s] %s", a.settings.Method, uri)
	}

	var reqBody io.Reader

	contentType := "application/json; charset=UTF-8"

	if a.settings.Method == methodPOST || a.settings.Method == methodPUT || a.settings.Method == methodPATCH {

		contentType = getContentType(input.Content)

		if input.Content != nil {
			if str, ok := input.Content.(string); ok {
				reqBody = bytes.NewBuffer([]byte(str))
			} else {
				b, _ := json.Marshal(input.Content) //todo handle error
				reqBody = bytes.NewBuffer([]byte(b))
			}
		}
	} else {
		reqBody = nil
	}

	req, err := http.NewRequest(a.settings.Method, uri, reqBody)

	if err != nil {
		return false, err
	}

	if reqBody != nil {
		req.Header.Set("Content-Type", contentType)
	}

	headers := input.Headers

	if len(headers) == 0 {
		headers = a.settings.Headers
	}

	// Set headers
	if len(headers) > 0 {
		if log.DebugEnabled() {
			log.Debug("Setting HTTP request headers...")
		}
		for key, value := range headers {
			if log.DebugEnabled() {
				log.Debugf("%s: %s", key, value)
			}
			req.Header.Set(key, value)
		}
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	log.Debug("response Status:", resp.Status)
	respBody, _ := ioutil.ReadAll(resp.Body)

	var result interface{}

	d := json.NewDecoder(bytes.NewReader(respBody))
	d.UseNumber()
	err = d.Decode(&result)

	//json.Unmarshal(respBody, &result)

	if log.DebugEnabled() {
		log.Debug("response body:", result)
	}

	output := &Output{Status: resp.StatusCode, Result: result}

	ctx.SetOutputObject(output)

	return true, nil
}

////////////////////////////////////////////////////////////////////////////////////////
// Utils

//todo just make contentType a setting
func getContentType(replyData interface{}) string {

	contentType := "application/json; charset=UTF-8"

	switch v := replyData.(type) {
	case string:
		if !strings.HasPrefix(v, "{") && !strings.HasPrefix(v, "[") {
			contentType = "text/plain; charset=UTF-8"
		}
	case int, int64, float64, bool, json.Number:
		contentType = "text/plain; charset=UTF-8"
	default:
		contentType = "application/json; charset=UTF-8"
	}

	return contentType
}

// BuildURI is a temporary crude URI builder
func BuildURI(uri string, values map[string]string) string {

	var buffer bytes.Buffer
	buffer.Grow(len(uri))

	addrStart := strings.Index(uri, "://")

	i := addrStart + 3

	for i < len(uri) {
		if uri[i] == '/' {
			break
		}
		i++
	}

	buffer.WriteString(uri[0:i])

	for i < len(uri) {
		if uri[i] == ':' {
			j := i + 1
			for j < len(uri) && uri[j] != '/' {
				j++
			}

			if i+1 == j {

				buffer.WriteByte(uri[i])
				i++
			} else {

				param := uri[i+1 : j]
				value := values[param]
				buffer.WriteString(value)
				if j < len(uri) {
					buffer.WriteString("/")
				}
				i = j + 1
			}

		} else {
			buffer.WriteByte(uri[i])
			i++
		}
	}

	return buffer.String()
}
