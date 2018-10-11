package rest

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Port int `md:"port,required"`
}

type HandlerSettings struct {
	Method string `md:"method,required,allowed(GET,POST,PUT,PATCH,DELETE)"`
	Path   string `md:"path,required"`
}

type Output struct {
	PathParams  map[string]string `md:"pathParams"`
	QueryParams map[string]string `md:"queryParams"`
	Headers     map[string]string `md:"headers"`
	Content     interface{}       `md:"content"`
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"pathParams":  o.PathParams,
		"queryParams": o.QueryParams,
		"headers":     o.Headers,
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
	o.Content = values["content"]

	return nil
}

type Reply struct {
	Code int         `md:"code"`
	Data interface{} `md:"data"`
}

func (r *Reply) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"code": r.Code,
		"data": r.Data,
	}
}

func (r *Reply) FromMap(values map[string]interface{}) error {

	var err error
	r.Code, err = coerce.ToInt(values["code"])
	if err != nil {
		return err
	}
	r.Data, _ = values["data"]

	return nil
}
