<!--
title: REST
weight: 4706
-->
# REST Trigger
This trigger provides your flogo application the ability to start an action via REST over HTTP

## Installation

```bash
flogo install github.com/project-flogo/contrib/trigger/rest
```

## Configuration

### Settings:
| Name      | Type   | Description
|:---       | :---   | :---       
| port      | int    | The port to listen on - **REQUIRED**
| enableTLS | bool   | Enable TLS on the server
| certFile  | string | The path to PEM encoded server certificate
| keyFile   | string | The path to PEM encoded server key


### Handler Settings:
| Name     | Type   | Description
|:---      | :---   | :---          
| method   | string | The HTTP method (ie. GET,POST,PUT,PATCH or DELETE) - **REQUIRED**
| path     | string | The resource path - **REQUIRED**

### Output:
| Name        | Type   | Description
|:---         | :---   | :---        
| pathParams  | params | The path parameters (e.g., 'id' in http://.../pet/:id/name )
| queryParams | params | The query parameters (e.g., 'id' in http://.../pet?id=someValue )
| headers     | params | The HTTP header parameters
| method      | string  | The HTTP method used for the request
| content     | any    | The content of the request

### Reply:
| Name    | Type   | Description
|:---     | :---   | :---        
| code    | int    | The http code to reply with
| data    | any    | The data to reply with
| headers | params | The HTTP response headers
| cookies | params | The HTTP response cookies to set (uses 'Set-Cookie' headers)

## Example Configurations

Triggers are configured via the triggers.json of your application. The following are some example configuration of the REST Trigger.

### POST
Configure the Trigger to handle a POST on /device

```json
{
  "triggers": [
    {
      "id": "flogo-rest",
      "ref": "github.com/project-flogo/contrib/trigger/rest",
      "settings": {
        "port": 8080
      },
      "handlers": [
        {
          "settings": {
            "method": "POST",
            "path": "/device"
          },
          "action": {
            "ref": "github.com/project-flogo/flow",
            "settings": {
              "flowURI": "res://flow:new_device_flow"
            }
          }
        }
      ]
    }
  ]
}
```

### GET
Configure the Trigger to handle a GET on /device/:id

```json
{
  "triggers": [
    {
      "id": "flogo-rest",
      "ref": "github.com/project-flogo/contrib/trigger/rest",
      "settings": {
        "port": 8080
      },
      "handlers": [
        {
          "settings": {
            "method": "GET",
            "path": "/device/:id"
          },
          "action": {
            "ref": "github.com/project-flogo/flow",
            "settings": {
              "flowURI": "res://flow:get_device_flow"
            },
            "input":{
              "deviceId":"=$.pathParams.id"
            }
          }
        }
      ]
    }
  ]
}
```
