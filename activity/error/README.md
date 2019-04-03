 <!-- 
title: Error
weight: 4610
-->

# Error
This activity allows you to cause an explicit error in the flow (throw an error).


## Installation

### Flogo CLI
```bash
flogo install  github.com/project-flogo/contrib/activity/error
```

## Configuration

### Input:
| Name     | Type   | Description
|:---      | :---   | :---    
| message  | string | The error message         
| data     | any    | The error data

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