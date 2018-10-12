<!--
title: REST
weight: 4618
-->

# REST
This activity allows you to invoke a REST service.

## Installation
### Flogo Web
This activity comes out of the box with the Flogo Web UI
### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/activity/rest
```

## Schema
Settings, Inputs and Outputs:

```json
{
  "settings":[
    {
      "name": "method",
      "type": "string",
      "required": true,
      "allowed" : ["GET", "POST", "PUT", "PATCH", "DELETE"]
    },
    {
      "name": "uri",
      "type": "string",
      "required": true
    },
    {
      "name": "proxy",
      "type": "string",
    },
    {
      "name": "headers",
      "type": "params"
    },
    {
      "name": "skipSSL",
      "type": "boolean",
      "value": "false"
    }
  ],
  "input":[
    {
      "name": "pathParams",
      "type": "params"
    },
    {
      "name": "queryParams",
      "type": "params"
    },
    {
      "name": "headers",
      "type": "params"
    },
    {
      "name": "content",
      "type": "any"
    }
  ],
  "output": [
    {
      "name": "result",
      "type": "any"
    },
    {
      "name": "status",
      "type": "int"
    }
  ]
}
```
## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| method      | true     | The HTTP method to invoke (Allowed values are GET, POST, PUT, DELETE, and PATCH) |         
| uri         | true     | The URI of the service to invoke |
| proxy       | false    | The address of the proxy server to be used |
| headers     | false    | The header parameters |
| skipSSL     | false    | If set to true, skips the SSL validation (defaults to false)

## Input
| Setting     | Required | Description |
|:------------|:---------|:------------|
| pathParams  | false    | The path parameters. This field is only required if you have params in your URI (for example http://.../pet/:id) |
| queryParams | false    | The query parameters |
| headers     | false    | The header parameters |
| content     | false    | The message content you want to send. This field is only used in POST, PUT, and PATCH |


## Examples
### Simple
The below example retrieves a pet with number '1234' from the [swagger petstore](http://petstore.swagger.io):

```json
{
  "id": "rest_2",
  "name": "Invoke REST Service",
  "description": "Simple REST Activity",
  "activity": {
    "ref": "github.com/TIBCOSoftware/flogo-contrib/activity/rest",
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
  "id": "rest_2",
  "name": "Rest 2",
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