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
| soapVersion             | string | SOAP Version to be used. Supported Version: 1.1 and 1.2 - **REQUIRED**
| timeout                 | integer| Request timeout in seconds
| enableTLS               | bool   | Set to true when using certificates. If set to false and SOAP Service Endpoint starts with https then certificate validation is set to false.
| xmlPassthroughMode      | bool   | Use this mode when SOAP body and headers are coming from upstream activity/trigger in XML format. When set, even incoming SOAP response is set in XML format without any processing. When not set, activity will convert request body and headers from JSON to XML and response body and headers from XML to JSON based on configured schema
| attributePrefix         | string | When xmlPassthroughMode is not set, use this confguration to define attribute representation in JSON data.  e.g If the prefix is set to '@', <car electric="true">Tesla</car> will be converted to {"car": {"@electric": "true", "#text": "Tesla"}}. The default is set to '-'
| serverCertificate       | string | When enableTLS set to true, you can configure a PEM encoded CA or Server certificate. You can set either base64 encoded certificate value or file path with prefix 'file://' 
| clientCertificate       | string | When enableTLS set to true, you can configure a PEM encoded Client certificate. You can set either base64 encoded certificate value or file path with prefix 'file://'
| clientKey               | string | When enableTLS set to true, you can configure a PEM encoded Clinent Key certificate. You can set either base64 encoded certificate value or file path with prefix 'file://'

### Input:
| Name               | Type   | Description
|:---                | :---   | :---     
| soapAction         | string | The SOAP action to be set in the header
| httpQueryParams    | any    | The HTTP query parameters to be sent. When xmlPassthroughMode set to false, you can coerce this field to JSON payload/schema. Just hover over the field and select 'Coerce with Schema' option after clicking '...'
| soapRequestHeaders | any    | The SOAP request headers to be sent. When xmlPassthroughMode set to false, you can coerce this field to JSON payload/schema. Just hover over the field and select 'Coerce with Schema' option after clicking '...'
| soapRequestBody    | any    |   The SOAP request headers to be sent. When xmlPassthroughMode set to false, you can coerce this field to JSON payload/schema. Just hover over the field and select 'Coerce with Schema' option after clicking '...'

### Output:
| Name                | Type   | Description
|:---                 | :---    | :---     
| httpStatus          | int     | The HTTP status code
| isFault             | bool    | Set to true when fault returned in the response
| soapResponsePayload | any     | The SOAP response payload received in case of success. When xmlPassthroughMode set to false, you can coerce this field to JSON payload/schema. Just hover over the field and select 'Coerce with Schema' option after clicking '...'
| soapResponseHeaders | any     | The SOAP response headers received. When xmlPassthroughMode set to false, you can coerce this field to JSON payload/schema. Just hover over the field and select 'Coerce with Schema' option after clicking '...'
| soapResponseFault   | any     | The SOAP response payload received in case of success. When xmlPassthroughMode set to false, you can coerce this field to JSON payload/schema. Just hover over the field and select 'Coerce with Schema' option after clicking '...'
