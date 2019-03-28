# Gateway with Javascript
This recipe is a gateway which runs some javascript.

## Installation
* Install [Go](https://golang.org/)
* Install the Flogo [CLI](https://github.com/project-flogo/cli)

## Setup
```bash
git clone https://github.com/project-flogo/jsexec
cd jsexec/examples/microgateway/json
```

## Testing
Create the gateway:
```bash
flogo create -f flogo.json
cd MyProxy
flogo build
```

Start the gateway:
```bash
bin/MyProxy
```

Run the following command:
```bash
curl http://localhost:9096/calculate"
```

You should see the following like response:
```json
{"sum":3}
```
