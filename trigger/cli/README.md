<!--
title: CLI
weight: 4701
-->
# CLI Trigger
This trigger provides your flogo application the ability to run as a CLI app, that is, accept input via the CLI & run once till completion and return the results to stdout.

## Installation

```bash
flogo install github.com/project-flogo/trigger/cli
```

## Configuration

###  Settings:
| Name      | Type   | Description
|:---       | :---   | :---     
| singleCmd | bool   | Indicates that this CLI runs only one command/handler         
| usage     | string | The usage details of the CLI
| long      | string | The description of the CLI

###  Handler Settings:
| Name  | Type   | Description
|:---   | :---   | :---
| flags | array  | List of flags         
| usage | string | The usage details of the command 
| short | string | A short description of the command
| long  | string | The description of the command

### Output:
| Name  | Type  | Description
|:---   | :---  | :---     
| args  | array | An array of the command line arguments  
| flags | map   | A map of the command line flags 

### Reply:
| Name | Type | Description
|:---  | :--- | :---     
| data | any  | The data that the command outputs |  


#### Flags
There is simple support for defining flags for a command.  You can specify either a boolean or string flag.
<br>
Flags are defined using the following format: `flagName||defaultValue||description`

_**Note:** if a flag has a default value of **true** or **false** it is considered a boolean flag_

## Sample Configuration
```json
"triggers": [
  {
    "id": "cli",
    "ref": "#cli",
    "name": "simple",
    "description": "Simple CLI Utility",
    "settings": {
      "singleCmd": true
    },
    "handlers": [
      {
        "name": "commandName",
        "settings": {
          "usage": "[flags] [args]",
          "short": "short command description",
          "long": "the long command descriptoin",
          "flags": [
           "flag1||||string flag",
           "flag2||false||boolan flag"
          ]
        },
        "action": {
          "ref": "#flow",
          "settings": {
            "flowURI": "res://flow:commandName"
          },
          "input": {
            "flags": "=$.flags",
            "args": "=$.args"
          }
        }
      }
    ]
  }
]  
```
_**Note:** Each CLI command maps to a handler, so in order to set your command a name, you must set the name of the handler._

# Examples

Triggers are configured via the triggers section of your application. The following are some example configuration of the CLI Trigger.

### Single command

An example can be found [here](examples/single).

### Multi command

An example can be found [here](examples/multi).

