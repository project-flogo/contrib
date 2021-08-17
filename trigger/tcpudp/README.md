<!--
title: TCP/UDP
weight: 4701
-->
# TCP/UDP Trigger

This trigger reads/writes data using TCI/UDP networks.

### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/trigger/tcpudp
```

## Configuration

### Setting :

| Name       | Type    | Description
|:---        | :---    | :---     
| network    | string  | Network type. Supported types: tcp,tcp4,tcp6,udp,udp4,udp6  - ***REQUIRED***
| host       | string  | Host IP or DNS resolvable name
| port       | string  | Port to listen on - ***REQUIRED***
| delimiter  | string  | Delimiter for read and write. If not set, trigger will take line delimiter '\n' as default value
| timeout    | integer | Read and Write timeout in milliseconds. To disable timeout, set value to 0.


### Output:

| Name         | Type     | Description
|:---          | :---     | :---   
| data         | string   | The data received from client

### Reply:

| Name         | Type     | Description
|:---          | :---     | :---   
| reply        | string   | The data to be sent back to the client

## Examples

```json
{
  "triggers": [
          {
              "ref": "github.com/project-flogo/contrib/trigger/tcpudp",
              "name": "ReceiveTCPData",
              "settings": {
                  "network": "tcp4",
                  "host": "localhost",
                  "port": "8999",
                  "delimiter": "\n",
                  "timeout": 200
              },
              "id": "ReceiveTCPData",
              "handlers": [
                  {
                      "settings": {},
                      "action": {
                          "ref": "github.com/project-flogo/flow",
                          "settings": {
                              "flowURI": "res://flow:TCP"
                          },
                          "input": {
                              "data": "=$.data"
                          },
                          "output": {
                              "reply": "=$.reply"
                          }
                      },
                      "reply": {
                          "reply": ""
                      }
                  }
              ]
          }
      ]
}
```