<!--
title: Timer
weight: 4707
-->
# Timer Trigger
This trigger provides your flogo application the ability to schedule an action

## Installation

```bash
flogo install github.com/project-flogo/contrib/trigger/timer
```

## Configuration

### Handler Settings:
| Name           | Type   | Description
|:---            | :---   | :---     
| startDelay     | string | The start delay (ex. 1m, 1h, etc.), immediate if not specified
| repeatInterval | string | The repeat interval (ex. 1m, 1h, etc.), doesn't repeat if not specified


## Example Configurations

Triggers are configured via the triggers.json of your application. The following are some example configuration of the Timer Trigger.

### Only once and immediate
Configure the Trigger to run a flow immediately

```json
{
  "triggers": [
    {
      "id": "flogo-timer",
      "ref": "github.com/project-flogo/contrib/trigger/timer",
      "handlers": [
        {
          "action": {
            "ref": "#flow",
            "settings": {
              "flowURI": "res://flow:myflow"
            }
          }
        }
      ]
    }
  ]
}
```

### Only once with a delay
Configure the Trigger to run a flow once with a delay of one minute.  "startDelay" settings format = "[hours]h[minutes]m[seconds]s"

```json
{
  "triggers": [
    {
      "id": "flogo-timer",
      "ref": "github.com/project-flogo/contrib/trigger/timer",
      "handlers": [
        {
          "settings": {
            "startDelay": "1m"
          },
          "action": {
            "ref": "#flow",
            "settings": {
              "flowURI": "res://flow:myflow"
            }
          }
        }
      ]
    }
  ]
}
```

### Repeating
Configure the Trigger to run a flow repeating every 10 minutes. "repeatInterval" settings format = "[hours]h[minutes]m[seconds]s"

```json
{
  "triggers": [
    {
      "id": "flogo-timer",
      "ref": "github.com/project-flogo/contrib/trigger/timer",
      "handlers": [
        {
          "settings": {
            "repeatInterval": "10m"
          },
          "action": {
            "ref": "github.com/project-flogo/flow",
            "settings": {
              "flowURI": "res://flow:myflow"
            }
          }
        }
      ]
    }
  ]
}
```

### Repeating with start delay
Configure the Trigger to run a flow every minute, with a delayed start of 10 minutes and 30 seconds.

```json
{
  "triggers": [
    {
      "id": "flogo-timer",
      "ref": "github.com/project-flogo/contrib/trigger/timer",
      "handlers": [
        {
          "settings": {
            "repeatInterval": "1m",
            "startDelay": "10m30s"
          },
          "action": {
            "ref": "github.com/project-flogo/flow",
            "settings": {
              "flowURI": "res://flow:myflow"
            }
          }
        }
      ]
    }
  ]
}
```
