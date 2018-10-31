<!--
title: REST
weight: 4706
-->
# REST Trigger
This trigger provides your flogo application the ability to start a flow via REST over HTTP

## Installation

```bash
flogo install github.com/project-flogo/contrib/trigger/rest
```

## Metadata
```json
{
  "settings": [
    {
      "name": "port",
      "type": "int",
      "required" : true
    }
  ],
  "handler": {
    "settings": [
      {
        "name": "method",
        "type": "string",
        "required" : true,
        "allowed" : ["GET", "POST", "PUT", "PATCH", "DELETE"]
      },
      {
        "name": "path",
        "type": "string",
        "required" : true
      }
    ]
  },
  "output": [
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
      "type": "object"
    }
  ],
  "reply": [
    {
      "name": "code",
      "type": "int"
    },
    {
      "name": "data",
      "type": "any"
    }
  ]
}
```
### Details
#### Trigger Settings:
| Setting  | Required | Description |
|:---------|:---------|:------------|
| port     | true     | The port to listen on


#### Handler Settings:
| Setting  | Required | Description |
|:---------|:---------|:------------|
| method   | true     | The HTTP method (ie. GET,POST,PUT,PATCH or DELETE)
| path     | true     | The resource path

#### Output:
|Name   | Description |
|:--------|:------------|
| pathParams  | The path params, ex. /device/:id, 'id' would be a path param
| queryParams | The query params
| headers     | The headers
| content     | The content of the request

#### Reply:
|Name   | Description |
|:--------|:------------|
| code  | The http code to reply with
| data  | The data to reply with


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
        "port": "8080"
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
        "port": "8080"
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
