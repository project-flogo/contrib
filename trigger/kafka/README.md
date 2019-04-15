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
| BrokerUrls | string | The Kafka cluster to connect to  |
| TrustStore | string | If connectiong to a TLS secured port, the directory containing the certificates representing the trust chain for the connection. This is usually just the CACert used to sign the server's certificate  |
| User       | string | If connectiong to a SASL enabled port, the userid to use for authentication  |
| Password   | string | If connectiong to a SASL enabled port, the password to use for authentication  |

### HandlerSettings:

| Name       | Type   | Description
|:---        | :---   | :---   
| Topic      | string | The Kafka topic on which to place the message  |
| Group      | string | The kafka group  |
| Partition  | string | Documents the partition that the message was placed on  |
| OffSet     | int64  | Documents the offset for the message  |

### Output:

| Name         | Type     | Description
|:---          | :---     | :---   
| Message      | string   | The text message sent  |


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