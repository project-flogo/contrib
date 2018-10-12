<!--
title: Return
weight: 4602
-->

# Return
This activity allows you to reply to a trigger invocation and map output values. After replying to the trigger, the flow ends (this will be the last actvity in your flow).

## Installation
### Flogo Web
This activity comes out of the box with the Flogo Web UI
### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/activity/reply
```

## Schema
Input and Output:

```json
{
  "settings":[
    {
      "name": "mappings",
      "type": "object",
      "required": true,
      "display": {
        "name": "Mapper",
        "type": "mapper",
        "mapperOutputScope" : "action.output"
      }
    }
  ],
  "input": [
  ],
  "output": [
  ]
}
```

## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| mappings    | true     | An set of mappings that are executed when the activity runs |


## Example
The below example allows you to configure the activity to reply and set the output values to literals "name" (a string) and 2 (an integer).

```json
{
  "id": "return",
  "name": "Return",
  "activity": {
    "ref": "github.com/project-flogo/contrib/activity/return",
    "settings": {
      "mappings": {
        "Output1":"name",
        "Output2":2
      }
    }
  }
}
```