package soapclient

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	SoapServiceEndpoint string `md:"soapServiceEndpoint,required"` // The URI of the service to invoke
	SoapVersion         string `md:"soapVersion"`                  // The address of the proxy server to be use
	Timeout             int    `md:"timeout"`                      // The request timeout in seconds
	EnableTLS           bool   `md:"enableTLS"`                    // Skip SSL validation
	XMLAttributePrefix  string `md:"xmlAttributePrefix"`           // The prefix to use for XML attributes while converting XML to JSON
	XMLMode             bool   `md:"xmlMode"`                      // Configure direct XML input/output
	ServerCertificate   string `md:"serverCertificate"`            // The server certificate
	ClientCertificate   string `md:"clientCertificate"`            // The client certificate
	ClientKey           string `md:"clientKey"`                    // The client key
}

type Input struct {
	SoapAction         string            `md:"soapAction"`
	HttpQueryParams    map[string]string `md:"httpQueryParams"`    // The HTTP query parameters
	SOAPRequestHeaders interface{}       `md:"soapRequestHeaders"` // The HTTP header parameters
	SOAPRequestBody    interface{}       `md:"soapRequestBody"`    // The message content to send. This is only used in POST, PUT, and PATCH
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"soapAction":         i.SoapAction,
		"httpQueryParams":    i.HttpQueryParams,
		"soapRequestHeaders": i.SOAPRequestHeaders,
		"soapRequestBody":    i.SOAPRequestBody,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {
	i.SOAPRequestHeaders = values["soapRequestHeaders"]
	i.SOAPRequestBody = values["soapRequestBody"]
	i.HttpQueryParams, _ = coerce.ToParams(values["httpQueryParams"])
	i.SoapAction, _ = values["soapAction"].(string)
	return nil
}

type Output struct {
	HttpStatus          int         `md:"httpStatus"`          // The HTTP status code
	IsFault             bool        `md:"isFault"`             // True if the response is a fault
	SOAPResponsePayload interface{} `md:"soapResponsePayload"` // The SOAP response content
	SOAPResponseHeaders interface{} `md:"soapResponseHeaders"` // The SOAP response headers
	SOAPResponseFault   interface{} `md:"soapResponseFault"`   // The SOAP fault content
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"httpStatus":          o.HttpStatus,
		"isFault":             o.IsFault,
		"soapResponsePayload": o.SOAPResponsePayload,
		"soapResponseHeaders": o.SOAPResponseHeaders,
		"soapResponseFault":   o.SOAPResponseFault,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {
	o.IsFault, _ = values["isFault"].(bool)
	o.SOAPResponsePayload = values["soapResponsePayload"]
	o.SOAPResponseHeaders = values["soapResponseHeaders"]
	o.SOAPResponseFault = values["soapResponseFault"]
	o.HttpStatus, _ = values["httpStatus"].(int)
	return nil
}
