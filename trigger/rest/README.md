<!--
title: REST
weight: 4706
-->
# REST Trigger
This trigger provides your flogo application the ability to start a flow via REST over HTTP

## Installation

```bash
flogo install github.com/TIBCOSoftware/flogo-contrib/trigger/rest
```

## Schema
Settings, Outputs and Endpoint:

```json
{
  "settings": [
    {
      "name": "port",
      "type": "integer",
      "required" : true
    }
  ],
  "handler": {
    "settings": [
      {
        "name": "method",
        "type": "string",
        "required" : true
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
  ]
}
```
## Settings
### Trigger:
| Setting  | Required | Description |
|:---------|:---------|:------------|
| port     | true     | The port to listen on


### Handler:
| Setting  | Required | Description |
|:---------|:---------|:------------|
| method   | true     | The HTTP method
| path     | true     | The resource path


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
