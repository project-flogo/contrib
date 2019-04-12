# Kafka Publisher Actvitiy

The `kafka` activity publishes the message on the kafka queue.

### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/activity/kafka
```

## Configuration


### Input:

| Name       | Type   | Description
|:---        | :---   | :---   
| BrokerUrls | string |   |
| Topic      | string |   |
| Message    | string |   |
| User       | string |   |
| Password   | string |   |
| TrustStore | string |   |

### Output:

| Name         | Type     | Description
|:---          | :---     | :---   
| Partition    | int32    |  |
| OffSet       | int64    |  |

## Examples

The below example sends `Hello From Flogo` to a Kafka Broker running on localhost:

```json
{
  "id": "publish_kafka_message",
  "name": "Publish Message to Kafka",
  "activity": {
    "ref": "github.com/project-flogo/contrib/activity/kafka",
    "input": {
      "brokerurls" : "localhost:9092",
      "topic"      : "syslog",
      "message"    : "Hello From Flogo",
      "user"       :  "",
      "password"   : ""
    }
  }
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