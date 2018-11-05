<!--
title: CLI
weight: 4701
-->
# CLI Trigger
This trigger provides your flogo application the ability to run as a CLI app, that is, accept input via the CLI & run once till completion and return the results to stdout.

## Installation

```bash
flogo install github.com/project-flogo/cli
```

## Metadata
```json
{
  "settings": [
    {
      "name": "singleCmd",
      "type": "bool"
    },
    {
      "name": "use",
      "type": "string"
    },
    {
      "name": "long",
      "type": "string"
    }
  ],
  "handler": {
    "settings": [
      {
        "name": "flags",
        "type": "array"
      },
      {
        "name": "use",
        "type": "string"
      },
      {
        "name": "short",
        "type": "string"
      },
      {
        "name": "long",
        "type": "string"
      }
    ]
  },
  "output": [
    {
      "name": "args",
      "type": "array"
    },
    {
      "name": "flags",
      "type": "object"
    }
  ],
  "reply": [
    {
      "name": "data",
      "type": "any"
    }
  ]
}
```
### Details
####  Settings:
| Setting      | Description                          |
|:-------------|:-------------------------------------|
| singleCmd    | Indicates that this cli runs only one command/handler |         
| use      | The usage details of the cli |
| long      | The description of the cli |
####  Handler Settings:
| Setting      | Description                          |
|:-------------|:-------------------------------------|
| flags      | The command invoked                  |         
| use      | The usage details of the  command |
| short      | A short description of the command |
| long      | The description of the command |

#### Output:
| Name      | Description                        |
|:------------|:-----------------------------------|
| args        | An array of the command line arguments |  
| flags       | A map of the command line flags |  

#### Reply:
| Name      | Description                        |
|:------------|:-----------------------------------|
| data        | The data that the command outputs |  


## Examples

Triggers are configured via the triggers section of your application. The following are some example configuration of the CLI Trigger.

### Single command

An example can be found [here](examples/single).

### Multi command

An example can be found [here](examples/multi).

