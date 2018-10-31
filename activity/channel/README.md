<!-- 
title: Channel
weight: 4603
-->

# Channel
This activity allows you to put a data on a named channel in the flogo engine.


## Installation
### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/activity/channel
```

## Metadata
```json
{
  "input":[
    {
      "name": "channel",
      "type": "string",
      "required": true
    },
    {
      "name": "data",
      "type": "interface{}",
      "required": true  
    }
  ]
}
```
### Details
#### Input:
| Name     | Required | Description |
|:------------|:---------|:------------|
| channel    | true     | The channel to put the value on |
| data    | true     | The data to put on the channel |

