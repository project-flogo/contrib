<!--
title: Log
weight: 4615
-->

# Log
This activity allows you to write log messages.

## Installation
### Flogo Web
This activity comes out of the box with the Flogo Web UI
### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/activity/log
```

## Schema
Inputs and Outputs:

```json
{
  "input":[
    {
      "name": "message",
      "type": "string",
      "value": ""
    },
    {
      "name": "addDetails",
      "type": "boolean",
      "value": "false"
    }
  ]
}
```
## Input
| Name     | Required | Description |
|:------------|:---------|:------------|
| message     | false    | The message to log |
| addDetails    | false    | If set to true this will append the  execution information to the log message |

## Examples
The below example logs a message 'test message':

```json
{
  "id": "log_message",
  "name": "Log Message",
  "activity": {
    "ref": "github.com/project-flogo/contrib/activity/log",
    "input": {
      "message": "test message",
      "addDetails": "false"
    }
  }
}
```