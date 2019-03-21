package rest

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Method        string            `md:"method,required,allowed(GET,POST,PUT,PATCH,DELETE)"`
	Uri           string            `md:"uri,required"`
	Headers       map[string]string `md:"headers"`
	Proxy         string            `md:"proxy"`
	SkipSSLVerify bool              `md:"skipSSL"`
	Timeout       int               `md:"timeout"`
	CertFile      string            `md:"certFile"`
	CAFile        string            `md:"caFile"`
	KeyFile       string            `md:"keyFile"`
}

type Input struct {
	PathParams  map[string]string `md:"pathParams"`
	QueryParams map[string]string `md:"queryParams"`
	Headers     map[string]string `md:"headers"`
	Content     interface{}       `md:"content"`
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
	Status int         `md:"status"`
	Data   interface{} `md:"result"`
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
