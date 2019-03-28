package rest

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Method        string            `md:"method,required,allowed(GET,POST,PUT,PATCH,DELETE)"` // The HTTP method to invoke
	Uri           string            `md:"uri,required"`  // The URI of the service to invoke
	Headers       map[string]string `md:"headers"`       // The HTTP header parameters
	Proxy         string            `md:"proxy"`         // The address of the proxy server to be use
	Timeout       int               `md:"timeout"`       // The request timeout in seconds
	SkipSSLVerify bool              `md:"skipSSLVerify"` // Skip SSL validation
	CertFile      string            `md:"certFile"`      // Path to PEM encoded client certificate
	KeyFile       string            `md:"keyFile"`       // Path to PEM encoded client key
	CAFile        string            `md:"CAFile"`        // Path to PEM encoded root certificates file

}

type Input struct {
	PathParams  map[string]string `md:"pathParams"`  // The query parameters (e.g., 'id' in http://.../pet?id=someValue )
	QueryParams map[string]string `md:"queryParams"` // The path parameters (e.g., 'id' in http://.../pet/:id/name )
	Headers     map[string]string `md:"headers"`     // The HTTP header parameters
	Content     interface{}       `md:"content"`     // The message content to send. This is only used in POST, PUT, and PATCH
}

func (o *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"pathParams":  o.PathParams,
		"queryParams": o.QueryParams,
		"headers":     o.Headers,
		"content":     o.Content,
	}
}

func (o *Input) FromMap(values map[string]interface{}) error {

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
	o.Content = values["content"]

	return nil
}

type Output struct {
	Status int         `md:"status"`  // The HTTP status code
	Data   interface{} `md:"result"`  // The HTTP response data
}

func (r *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"status": r.Status,
		"data":   r.Data,
	}
}

func (r *Output) FromMap(values map[string]interface{}) error {

	var err error
	r.Status, err = coerce.ToInt(values["status"])
	if err != nil {
		return err
	}
	r.Data, _ = values["data"]

	return nil
}
