package rest

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Port      int    `md:"port,required"` // The port to listen on
	EnableTLS bool   `md:"enableTLS"`     // Enable TLS on the server
	CertFile  string `md:"certFile"`      // The path to PEM encoded server certificate
	KeyFile   string `md:"keyFile"`       // The path to PEM encoded server key
}

type HandlerSettings struct {
	Method string `md:"method,required,allowed(GET,POST,PUT,PATCH,DELETE)"` // The HTTP method (ie. GET,POST,PUT,PATCH or DELETE)
	Path   string `md:"path,required"`                                      // The resource path
}

type Output struct {
	PathParams  map[string]string `md:"pathParams"`  // The path parameters (e.g., 'id' in http://.../pet/:id/name )
	QueryParams map[string]string `md:"queryParams"` // The query parameters (e.g., 'id' in http://.../pet?id=someValue )
	Headers     map[string]string `md:"headers"`     // The HTTP header parameters
	Content     interface{}       `md:"content"`     // The content of the request
	Method      string            `md:"method"`      // The HTTP method used for the request
}

type Reply struct {
	Code    int               `md:"code"`    // The http code to reply with
	Data    interface{}       `md:"data"`    // The data to reply with
	Headers map[string]string `md:"headers"` // The HTTP response headers
	Cookies []interface{}     `md:"cookies"` // "The response cookies, adds `Set-Cookie` headers"
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"pathParams":  o.PathParams,
		"queryParams": o.QueryParams,
		"headers":     o.Headers,
		"method":      o.Method,
		"content":     o.Content,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.PathParams, err = coerce.ToParams(values["pathParams"])
	if err != nil {
		return err
	}
	o.QueryParams, err = coerce.ToParams(values["queryParams"])
	if err != nil {
		return err
	}
	o.Headers, err = coerce.ToParams(values["headers"])
	if err != nil {
		return err
	}
	o.Method, err = coerce.ToString(values["method"])
	if err != nil {
		return err
	}
	o.Content = values["content"]

	return nil
}

func (r *Reply) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"code":    r.Code,
		"data":    r.Data,
		"headers": r.Headers,
		"cookies": r.Cookies,
	}
}

func (r *Reply) FromMap(values map[string]interface{}) error {

	var err error
	r.Code, err = coerce.ToInt(values["code"])
	if err != nil {
		return err
	}
	r.Data, _ = values["data"]

	r.Headers, err = coerce.ToParams(values["headers"])
	if err != nil {
		return err
	}

	r.Cookies, err = coerce.ToArray(values["cookies"])
	if err != nil {
		return err
	}

	return nil
}
