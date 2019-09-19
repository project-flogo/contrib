<!--
title: JSExec
weight: 4616
-->

# Javascript Execution Activity

The `jsexec` activity evaluates a javascript `script` along with provided `parameters` and returns the result in the `outputs`.

### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/activity/jsexec
```

## Configuration

### Settings:

| Name   | Type   | Description
|:---    | :---   | :---   
| script | string | The javascript code to evaluate |

### Input:

| Name       | Type   | Description
|:---        | :---   | :---   
| parameters | object | Key/value pairs representing parameters to evaluate within the context of the script  |

### Output:

| Name         | Type     | Description
|:---          | :---     | :---   
| error        | bool   | Flag indicating if there was an error executing the script |
| errorMessage | string | The error message |
| result       | object | The result object from the javascript code  |

## Microgateway Usage

A sample `service` definition is:

```json
{
  "name": "JSCalc",
  "description": "Make calls to a JS calculator",
  "ref": "github.com/project-flogo/jsexec",
  "settings": {
    "script": "result.total = parameters.num * 2;"
  }
}
```

An example `step` that invokes the above `JSCalc` service using `parameters` is:

```json
{
  "if": "$.PetStorePets.outputs.result.status == 'available'",
  "service": "JSCalc",
  "input": {
    "parameters.num": "=$.PetStorePets.outputs.result.available"
  }
}
```

Utilizing the response values can be seen in a response handler:

```json
{
  "if": "$.PetStorePets.outputs.result.status == 'available'",
  "error": false,
  "output": {
    "code": 200,
    "data": {
      "body.pet": "=$.PetStorePets.outputs.result",
      "body.inventory": "=$.PetStoreInventory.outputs.result",
      "body.availableTimesTwo": "=$.JSCalc.outputs.result.total"
    }
  }
}
```
Additional microgateway examples that utilize the `jsexec` activity

* [api](examples/microgateway/api/README.md) - A simple API example
* [json](examples/microgateway/api/README.md) - An example that using a  _flogo.json_
 
## Development

### Testing

To run tests issue the following command in the root of the project:

```bash
go test -p 1 ./...
```

The `-p 1` is needed to prevent tests from being run in parallel. To re-run the tests first run the following:

```bash
go clean -testcache
```

To skip the integration tests use the `-short` flag:

```bash
go test -p 1 -short ./...
```