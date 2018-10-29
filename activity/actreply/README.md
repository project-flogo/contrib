<!--
title: Reply
weight: 4601
-->

# Reply
This activity allows you to reply to a trigger invocation and map output values. After replying to the trigger, this activity will allow the action to continue further.

## Installation
### Flogo Web
This activity comes out of the box with the Flogo Web UI
### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/activity/actreply
```

## Metadata

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
  "id": "reply",
  "name": "Reply",
  "description": "Simple Reply Activity",
  "activity": {
    "ref": "github.com/project-flogo/contrib/activity/actreply",
    "settings": {
      "mappings": {
        "Output1":"name",
        "Output2":2
      }
    }
  }
}
```