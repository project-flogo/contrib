# RunAction
This activity allows you to run flogo actions.

## Installation

### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/activity/runaction
```

## Configuration

### Settings
| Name          | Type   | Description
|:---           | :---   | :---    
| actionRef     | string | The 'ref' to the action to be run
| actionSettings| object | The settings of the action

### Input
The inputs for this activity should be the inputs for the action you are running

### Output
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
