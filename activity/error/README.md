 <!-- 
title: Error
weight: 4610
-->

# Error
This activity allows you to cause an explicit error in the flow (throw an error).


## Installation
### Flogo Web
This activity comes out of the box with the Flogo Web UI
### Flogo CLI
```bash
flogo install  github.com/project-flogo/contrib/activity/error
```

## Metadata
```json
{
  "input":[
    {
      "name": "message",
      "type": "string"
    },
    {
      "name": "data",
      "type": "object"
    }
  ]
}
```
### Details
#### Input:
| Name     | Required | Description |
|:------------|:---------|:------------|
| message     | false    | The error message you want to throw |         
| data        | false    | The error data you want to throw |

## Configuration Examples
The below example throws a simple error with a message:

```json
{
  "id": "throw_error",
  "name": "Throw Error",
  "activity": {
    "ref": "github.com/project-flogo/contrib/activity/error",
    "input": {
      "message": "Unexpected Threshold Value"
    }
  }
}
```