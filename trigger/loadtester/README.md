<!--
title: LoadTester
weight: 4706
-->
# LoadTester Trigger
This trigger provides your flogo application the ability to run simple load test on a specified action

Implementation based off [go-wrk](github.com/tsliwowicz/go-wrk).

## Installation

```bash
flogo install github.com/project-flogo/contrib/trigger/loadtester
```

## Configuration    

###  Settings:
| Name        | Type   | Description
|:---         | :---   | :---     
| startDelay  | int    | The start delay of the test in seconds, default: 30
| duration    | int    | The duration of the test in seconds, default: 60
| concurrency | int    | The level of concurrency, default: 5
| data        | any    | Optional data to pass along to the action
| handler     | string | The named handler to test, defaults to the first handler

#### Output:
| Name  | Type | Description
|:---   | :--- | :---     
| data  | any  | The data from the settings to pass along

## Example Configuration

### Test Flow
Configure the Trigger to load test the 'flow:testflow'

```json
{
  "triggers": [
    {
      "id": "tester",
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
            "ref": "github.com/project-flogo/flow",
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