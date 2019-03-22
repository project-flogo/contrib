# Gateway with Javascript
This recipe is a gateway which runs some javascript.

## Installation
* Install [Go](https://golang.org/)

## Setup
```bash
git clone https://github.com/project-flogo/contrib/activity/jsexec
cd jsexec/examples/microgateway/api
```

## Testing

Start the gateway:
```bash
go run example.go
```

Run the following command:
```bash
curl http://localhost:9096/calculate"
```

You should see the following like response:
```json
{"sum":3}
```
