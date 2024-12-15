<!--
title: SOAP Client
weight: 4618
-->

# SOAP Client
This activity allows you to invoke a SOAP service.

## Installation

### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/activity/soapclient
```

## Configuration

### Settings:
| Name                    | Type   | Description
|:---                     | :---   | :---     
| soapServiceEndpoint     | string | SOAP Service Endpoint - **REQUIRED**
| soapVersion             | string | SOAP Version to be used. Supported Versions: **1.1** and **1.2** - **REQUIRED**
| timeout                 | integer| Request timeout in seconds
| enableTLS               | bool   | Set to true when using certificates. If set to false and SOAP Service Endpoint starts with `https` then certificate validation is set to false.
| xmlMode                 | bool   | Use this mode when you want to directly map XML data to request SOAP header and body and access SOAP response in XML data format. When not set, activity will convert configured request body and headers from JSON to XML and response body and headers from XML to JSON based on JSON schema
| xmlAttributePrefix      | string | When `xmlMode=false`, use this confguration to define XML attribute representation in JSON data. e.g If prefix is set to '@', `<car electric="true">Tesla</car>` will be converted to `{"car": {"@electric": "true", "#text": "Tesla"}}`. The default is set to '-'
| serverCertificate       | string | When enableTLS set to true, configure a PEM encoded CA or Server certificate. Set either base64 encoded certificate value or file path with prefix `file://` 
| clientCertificate       | string | When enableTLS set to true, configure a PEM encoded Client certificate. Set either base64 encoded certificate value or file path with prefix `file://`
| clientKey               | string | When enableTLS set to true, configure a PEM encoded Clinent Key certificate. Set either base64 encoded certificate value or file path with prefix `file://`

### Input:
| Name               | Type   | Description
|:---                | :---   | :---     
| soapAction         | string | The SOAP action to be set in the header
| httpQueryParams    | any    | The HTTP query parameters to be sent. You can configure JSON schema for this field. To configure schema, hover over the field and select `Coerce with Schema` option after clicking `...`
| soapRequestHeaders | any    | The SOAP request headers to be sent. When `xmlMode=false`, you can configure JSON schema for this field. To configure schema, hover over the field and select `Coerce with Schema` option after clicking `...`
| soapRequestBody    | any    |   The SOAP request headers to be sent. When `xmlMode=false`, you can configure JSON schema for this field. To configure schema, hover over the field and select `Coerce with Schema` option after clicking `...`

### Output:
| Name                | Type   | Description
|:---                 | :---    | :---     
| httpStatus          | int     | The HTTP status code
| isFault             | bool    | Set to true when fault returned in the response
| soapResponsePayload | any     | The SOAP response payload received in case of success. When `xmlMode=false`, you can configure JSON schema for this field. To configure schema, hover over the field and select `Coerce with Schema` option after clicking `...`. When `xmlMode=true`, output will be XML string.
| soapResponseHeaders | any     | The SOAP response headers received. When `xmlMode=false`, you can configure JSON schema for this field. To configure schema, hover over the field and select `Coerce with Schema` option after clicking `...`. When `xmlMode=true`, output will be XML string.
| soapResponseFault   | any     | The SOAP response payload received in case of success. When `xmlMode=false`, you can configure JSON schema for this field. To configure schema, hover over the field and select `Coerce with Schema` option after clicking `...`. When `xmlMode=true`, output will be XML string. 
