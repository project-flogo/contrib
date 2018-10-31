<!--
title: Channel
weight: 4706
-->
# Channel Trigger
This trigger provides your flogo application the ability to start an action via a named engine channel

## Installation

```bash
flogo install github.com/project-flogo/contrib/trigger/channel
```

## Metadata
```json
{
  "handler": {
    "settings": [
      {
        "name": "channel",
        "type": "string",
        "required" : true
      }
    ]
  },
  "output": [
    {
      "name": "data",
      "type": "any"
    }
  ]
}
```
### Details    
#### Handler Settings:
| Setting  | Required | Description |
|:---------|:---------|:------------|
| channel  | true     | The internal engine channel |

#### Output:
|Name   | Description |
|:--------|:------------|
| data     | The data pulled from the channel


## Example Configurations

Triggers are configured via the triggers.json of your application. The following are some example configuration of the Channel Trigger.

### Run Flow
Configure the Trigger to handle an event recieved on the 'test' channel

```json
{
  "triggers": [
    {
      "id": "flogo-channel",
      "ref": "github.com/project-flogo/contrib/trigger/channel",
      "handlers": [
        {
          "settings": {
            "channel": "test"
          },
          "action": {
            "ref": "github.com/TIBCOSoftware/flogo-contrib/action/flow",
            "settings": {
                "flowURI": "res://flow:testflow"
            }       
          }
        }
      ]
    }
  ]
}
```

Note: a channel must be defined in the app in order to use it in a trigger or activity.  A channel is specified by "\<name\>:\<buffer size\>"
```
"channels": [
  "test:5"
]
```