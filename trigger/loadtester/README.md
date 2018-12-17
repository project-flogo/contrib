<!--
title: LoadTester
weight: 4706
-->
# LoadTester Trigger
This trigger provides your flogo application the ability to run simple load test on a specified action

Implementation based off github.com/tsliwowicz/go-wrk

## Installation

```bash
flogo install github.com/project-flogo/contrib/trigger/loadtester
```

## Metadata
```json
{
  "settings": [
    {
      "name": "startDelay",
      "type": "int"
    },
    {
      "name": "duration",
      "type": "int"
    },
    {
      "name": "concurrency",
      "type": "int"
    },
    {
      "name": "data",
      "type": "any"
    },
    {
      "name": "handler",
      "type": "string"
    }
  ],
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
| startDelay  | false | The start delay of the test in seconds, default: 30|
| duration    | false | The duration of the test in seconds, default: 60 |
| concurrency | false | The level of concurrency, default: 5 |
| data        | false | Optional data to pass along to the action |
| handler     | true  | The named handler to test, defaults to the first handler |

#### Output:
|Name   | Description |
|:--------|:------------|
| data     | The data from the settings to pass along |

## Example Configuration

### Test Flow
Configure the Trigger to load test the 'flow:testflow'

```json
{
  "triggers": [
    {
      "id": "flogo-channel",
      "ref": "github.com/project-flogo/contrib/trigger/loadtester",
      "settings": {
        "startDelay": 15,
        "duration": 120,
        "concurrency" : 5,
        "handler": "test"
      },
      "handlers": [
        {
          "name": "test",
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
`````