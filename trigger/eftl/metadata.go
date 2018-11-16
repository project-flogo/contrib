package eftl

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	URL 		string `md:"url"`
	Id 		string `md:"id"`
	User 	 	string `md:"user"`
	Password 	string `md:"password"`
	CA 		string `md:"ca"`
	Tracer 		string `md:"tracer"`
	TracerEndpoint 	string `md:"tracerEndpoint"`
	TracerToken 	string `md:"tracerToken"`
	TracerDebug 	string `md:"tracerDebug"`
	TracerSameSpan 	string `md:"tracerSameSpan"`
	TracerId128bit 	string `md:"tracerID128Bit"`
}

type HandlerSettings struct {
	Dest 		string `md:"dest"`
	Condition   	string `md:"condition"`
}

type Output struct {
	PathParams	map[string]string 	`md:"pathParams"`
	QueryParams 	map[string]string 	`md:"queryParams"`
	Params      	map[string]string 	`md:"params"`
	Content     	interface{}       	`md:"content"`
	//Tracing	  	map[string]interface{} 	`md:"tracing"`

}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"pathParams":  o.PathParams,
		"queryParams": o.QueryParams,
		"params":     o.Params,
		"content":     o.Content,
		//"tracing":	o.Tracing,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {
	o.PathParams = values["pathParams"].(map[string]string)
	o.QueryParams = values["queryParams"].(map[string]string)
	o.Params = values["params"].(map[string]string)
	o.Content = values["content"].(interface{})
	//o.Tracing = values["tracing"].(map[string]interface{})
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