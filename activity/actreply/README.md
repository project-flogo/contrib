<!--
title: Reply
weight: 4601
-->

# Reply
This activity allows you to reply to a trigger invocation and map output values. After replying to the trigger, this activity will allow the action to continue further.

## Installation

### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/activity/actreply
```

## Configuration

### Settings:
| Name     | Type   | Description
|:---      | :---   | :---    
| mappings | object | Set of mappings to execute when the activity runs

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