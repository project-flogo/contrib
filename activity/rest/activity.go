package rest

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/ssl"
)

func init() {
	_ = activity.Register(&Activity{}, New)
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

	client := &http.Client{}

	httpTransportSettings := &http.Transport{}

	if s.Timeout > 0 {
		httpTransportSettings.ResponseHeaderTimeout = time.Second * time.Duration(s.Timeout)
	}

	logger := ctx.Logger()

	// Set the proxy server to use, if supplied
	if len(s.Proxy) > 0 {
		proxyURL, err := url.Parse(s.Proxy)
		if err != nil {
			logger.Debugf("Error parsing proxy url '%s': %s", s.Proxy, err)
			return nil, err
		}

		logger.Debug("Setting proxy server:", s.Proxy)
		httpTransportSettings.Proxy = http.ProxyURL(proxyURL)
	}

	if strings.HasPrefix(s.Uri, "https") {

		cfg := &ssl.Config{}

		if len(s.SSLConfig) != 0 {
			err := cfg.FromMap(s.SSLConfig)
			if err != nil {
				return nil, err
			}

			if _, set := s.SSLConfig["skipVerify"]; !set {
				cfg.SkipVerify = true
			}
			if _, set := s.SSLConfig["useSystemCert"]; !set {
				cfg.UseSystemCert = true
			}
		} else {
			//using ssl but not configured, use defaults
			cfg.SkipVerify = true
			cfg.UseSystemCert = true
		}

		tlsConfig, err := ssl.NewClientTLSConfig(cfg)
		if err != nil {
			return nil, err
		}

		httpTransportSettings.TLSClientConfig = tlsConfig
	}

	client.Transport = httpTransportSettings
	act.client = client

	return act, nil
}

// Activity is an activity that is used to invoke a REST Operation
// settings : {method, uri, headers, proxy, skipSSL}
// input    : {pathParams, queryParams, headers, content}
// outputs  : {status, result}
type Activity struct {
	settings      *Settings
	containsParam bool
	client        *http.Client
}

func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Invokes a REST Operation
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, err
	}

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

	logger := ctx.Logger()

	if logger.DebugEnabled() {
		logger.Debugf("REST Call: [%s] %s", a.settings.Method, uri)
	}

	var reqBody io.Reader

	contentType := "application/json; charset=UTF-8"
	method := a.settings.Method

	if method == methodPOST || method == methodPUT || method == methodPATCH {

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

	req, err := http.NewRequest(method, uri, reqBody)
	if err != nil {
		return false, err
	}

	if reqBody != nil {
		req.Header.Set("Content-Type", contentType)
	}

	headers := a.getHeaders(input.Headers)

	// Set headers
	if len(headers) > 0 {
		if logger.DebugEnabled() {
			logger.Debug("Setting HTTP request headers...")
		}
		for key, value := range headers {
			if logger.TraceEnabled() {
				logger.Trace("%s: %s", key, value)
			}
			req.Header.Set(key, value)
		}
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return false, err
	}

	if resp == nil {
		logger.Trace("Empty response")
		return true, nil
	}

	defer func() {
		if resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()

	if logger.DebugEnabled() {
		logger.Debug("Response status:", resp.Status)
	}

	respHeaders := make(map[string]string, len(resp.Header))

	for key := range resp.Header {
		respHeaders[key] = resp.Header.Get(key)
	}

	var cookies []interface{}

	for _, cookie := range resp.Header["Set-Cookie"] {
		cookies = append(cookies, cookie)
	}

	var result interface{}

	// Check the HTTP Header Content-Type
	respContentType := resp.Header.Get("Content-Type")
	switch respContentType {
	case "application/json":
		d := json.NewDecoder(resp.Body)
		d.UseNumber()
		err = d.Decode(&result)
		if err != nil {
			switch {
			case err == io.EOF:
				// empty body
			default:
				return false, err
			}
		}
	default:
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return false, err
		}

		result = string(b)
	}

	if logger.TraceEnabled() {
		logger.Trace("Response body:", result)
	}


	output := &Output{Status: resp.StatusCode, Data: result, Headers:respHeaders, Cookies:cookies}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}

////////////////////////////////////////////////////////////////////////////////////////
// Utils

func (a *Activity) getHeaders(inputHeaders map[string]string) map[string]string {

	if len(inputHeaders) == 0 {
		return a.settings.Headers
	}

	if len(a.settings.Headers) == 0 {
		return inputHeaders
	}

	headers := make(map[string]string)
	for key, value := range a.settings.Headers {
		headers[key] = value
	}
	for key, value := range inputHeaders {
		headers[key] = value
	}

	return headers
}


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
