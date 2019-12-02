<!--
title: REST
weight: 4618
-->

# REST
This activity allows you to invoke a REST service.

## Installation

### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/activity/rest
```

## Configuration

### Settings:
| Name          | Type   | Description
|:---           | :---   | :---     
| method        | string | The HTTP method to invoke (Allowed values are GET, POST, PUT, DELETE, and PATCH) - **REQUIRED**
| uri           | string | The URI of the service to invoke - **REQUIRED**
| headers       | params | The HTTP header parameters
| proxy         | string | The address of the proxy server to be used
| timeout       | int    | The request timeout in seconds
| sslConfig     | object | SSL configuration

#### *sslConfig* Object: 
| Property      | Type   | Description
|:---           | :---   | :---     
| skipVerify    | bool   | Skip SSL validation, defaults to true
| useSystemCert | bool   | Use the systems root certificate file, defaults to true
| caFile        | string | The path to PEM encoded root certificates file
| certFile      | string | The path to PEM encoded client certificate
| keyFile       | string | The path to PEM encoded client key

*Note: used if URI is https*
### Input:
| Name        | Type   | Description
|:---         | :---   | :---     
| pathParams  | params | The path parameters (e.g., 'id' in http://.../pet/:id/name )
| queryParams | params | The query parameters (e.g., 'id' in http://.../pet?id=someValue )
| headers     | params | The HTTP header parameters
| content     | any    | The message content to send. This is only used in POST, PUT, and PATCH

### Output:
| Name    | Type   | Description
|:---     | :---   | :---     
| status  | int    | The HTTP status code
| data    | any    | The HTTP response data
| headers | params | The HTTP response headers
| cookies | array  | The response cookies (from 'Set-Cookie')

## Examples
### Simple
The below example retrieves a pet with number '1234' from the [swagger petstore](http://petstore.swagger.io):

```json
{
  "id": "rest_activity",
  "name": "REST Activity",
  "activity": {
    "ref": "github.com/project-flogo/contrib/activity/rest",
    "settings": {
      "method": "GET",
      "uri": "http://petstore.swagger.io/v2/pet/1234"
    }
  }
}
```

### Using Path Params
The below example is the same as above, it retrieves a pet with number '1234' from the [swagger petstore](http://petstore.swagger.io), but uses a URI parameter to configure the ID:

```json
{
  "id": "rest_activity",
  "name": "REST Activity",
  "activity": {
    "ref": "github.com/project-flogo/contrib/activity/rest",
    "settings": {
      "method": "GET",
      "uri": "http://petstore.swagger.io/v2/pet/:id"
    },
    "input": {
      "params": { "id": "1234"}
    }
  }
}
```