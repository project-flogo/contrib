# Run Action
This activity allows you to run flogo actions.

## Installation

### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/activity/runaction
```

## Configuration

### Setting:
| Name          | Type   | Description
|:---           | :---   | :---    
| actionRef     | string | The 'ref' to the type of flogo action.
| actionSettings| object | The settings object of the flogo action. 

### Input:
| Name          | Type   | Description
|:---           | :---   | :---    
// Note : The Inputs should be inputs for the action

### Output:
| Name          | Type   | Description
|:---           | :---   | :---    
| output        | object | The output of the action.


## Examples
The below example logs a message 'test message':

```json
{
    "id": "cmlact",
    "ref": "github.com/project-flogo/contrib/activity/runaction",
    "settings": {
        "actionRef": "github.com/project-flogo/catalystml-flogo/action",
        "actionSettings": { "catalystMlURI" : "file://cml.json" }
    },
    "input": {
        "dataIn": "=$.result"
    }
}          
```
