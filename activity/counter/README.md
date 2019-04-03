<!-- 
title: Counter
weight: 4609
-->

# Counter
This activity allows you to use a global counter.

## Installation

### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/activity/counter
```

## Configuration

### Settings:
| Name        | Type   | Description
|:---         | :---   | :---    
| counterName | string | The name of the counter - **REQUIRED**         
| op          | string | The counter operation, 'get' is the default operation

### Output:
| Name  | Type | Description
|:---   | :--- | :---    
| value | int  |  The result of the counter operation

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