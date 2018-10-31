<!-- 
title: Counter
weight: 4609
-->

# Counter
This activity allows you to use a global counter.

## Installation
### Flogo Web
This activity comes out of the box with the Flogo Web UI
### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/activity/counter
```

## Metadata
```json
{
  "settings":[
    {
      "name": "counterName",
      "type": "string",
      "required": true
    },
    {
      "name": "op",
      "type": "string",
      "allowed" : ["get", "increment", "reset"]
    }
  ],
  "output": [
    {
      "name": "value",
      "type": "integer"
    }
  ]
}
```
### Details
#### Settings:
| Setting     | Required | Description |
|:------------|:---------|:------------|
| counterName | true     | The name of the counter |         
| op          | false    | Counter operation, 'get' is the default operation|

#### Output:
|Name   | Description |
|:--------|:------------|
| value  | the result of the counter operation

## Examples
### Increment
The below example increments a 'messages' counter:

```json
{
  "id": "increment_message_count",
  "name": "Increment Message Count",
  "activity": {
    "ref": "github.com/project-flogo/contrib/activity/counter",
    "settings": {
      "counterName": "messages",
      "op": "increment"
    }
  }
}
```

### Get
The below example retrieves the last value of the 'messages' counter:

```json
{
  "id": "get_message_count",
  "name": "Get Message Count",
  "activity": {
    "ref": "github.com/project-flogo/contrib/activity/counter",
    "settings": {
      "counterName": "messages",
      "op": "get"
    }
  }
}
```

### Reset
The below example resets the 'messages' counter:

```json
{
  "id": "reset_message_count",
  "name": "Reset Message Count",
  "activity": {
    "ref": "github.com/project-flogo/contrib/activity/counter",
    "settings": {
      "counterName": "messages",
      "op": "reset"
    }
  }
}
```