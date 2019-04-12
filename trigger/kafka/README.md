# Kafka Subscriber Trigger

The `kafka` trigger subscribes to a topic on kafka broker and listens for the messages.

### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/trigger/kafka
```

## Configuration

### Setting :

| Name       | Type   | Description
|:---        | :---   | :---     
| BrokerUrls | string |   |
| TrustStore | string |   |
| User       | string |   |
| Password   | string |   |

### HandlerSettings:

| Name       | Type   | Description
|:---        | :---   | :---   
| Topic      | string |   |
| Group      | string |   |
| Partition  | string |   |
| OffSet     | int64  |   |

### Output:

| Name         | Type     | Description
|:---          | :---     | :---   
| Message      | string   |  |


## Examples

```json
{
  "triggers": [
    {
      "id": "flogo-kafka",
      "ref": "github.com/project-flogo/contrib/trigger/kafka",
      "settings": {
        "brokerurls" : "localhost:9092",
        "truststore" : "" 
      },
      "handlers": [
        {
          "settings": {
            "topic": "syslog",
            "group": ""
          },
          "action": {
            "ref": "github.com/project-flogo/flow",
            "settings": {
              "flowURI": "res://flow:new_device_flow"
            }
          }
        }
      ]
    }
  ]
}
```
 
## Development

### Testing

To run tests first set up the kafka broker using the docker-compose file given below:

```yaml
version: '2'
  
services:

  zookeeper:
    image: wurstmeister/zookeeper:3.4.6
    expose:
    - "2181"

  kafka:
    image: wurstmeister/kafka:2.11-2.0.0
    depends_on:
    - zookeeper
    ports:
    - "9092:9092"
    environment:
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092
      KAFKA_LISTENERS: PLAINTEXT://0.0.0.0:9092
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
```

Then run the following command: 

```bash
go test 
```