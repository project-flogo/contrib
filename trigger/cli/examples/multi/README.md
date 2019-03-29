
# Multi-Command Example
This example shows how to configure the CLI trigger to run as a cli with multiple commands. Help and version commands are automatically generated for multi-command CLIs.

To build and run cli example and execute
```
flogo create -f flogo-multi-cli.json
cd cli
flogo build --shim cli
./bin/cli
```


## Configuration
```json
{
 "triggers": [
     {
       "id": "cli",
       "ref": "#cli",
       "name": "simple",
       "description": "Simple CLI Utility",
       "settings": {
         "usage":"<command>",
         "long":"A simple cli using flogo"
       },
       "handlers": [
         {
           "name":"test1",
           "settings": {
             "usage":"[flags] [args]",
             "short": "test command",
             "long": "the test command",
             "flags": [
               "flag1||defaultValue||the first value flag",
               "flag2||false||the first bool flag"
             ]
           },
           "action": {
             "ref": "#flow",
             "settings": {
               "flowURI": "res://flow:command1"
             },
             "input" :{
               "flags":"=$.flags",
               "args":"=$.args"
             }
           }
         },
         {
           "name":"test2",
           "settings": {
             "usage":"[flags] [args]",
             "short": "test2 command",
             "long": "the test2 command",
             "flags": [
               "flag1||defaultValue||the first value flag",
               "flag2||false||the first bool flag"
             ]
           },
           "action": {
             "ref": "#flow",
             "settings": {
               "flowURI": "res://flow:command2"
             },
             "input" :{
               "flags":"=$.flags",
               "args":"=$.args"
             }
           }
         }
       ]
     }
   ]
 }
```


### Help
General Help: `cli help`
```
A simple cli using flogo
Usage:
    cli <command>

Commands:
    test1        test command
    test2        test2 command
    help         help on command
    version      prints cli version
```

Command Help: `cli help test1`

```
the test command
Usage:
   cli test1 [flags] [args]

Flags: 
   -flag1 string        the first value flag
   -flag2               the first bool flag
```
