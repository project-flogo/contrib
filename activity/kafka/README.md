<!--
title: Kafka
weight: 4616
-->

# Kafka Activity

This activity publishes messages on a topic in a Kafka cluster.

### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/activity/kafka
```

## Configuration

### Settings: 
| Name       | Type   | Description
|:---        | :---   | :---   
| brokerUrls | string | The brokers of the Kafka cluster to connect to - ***REQUIRED***
| topic      | string | The Kafka topic on which to place the message - ***REQUIRED***
| user       | string | If connecting to a SASL enabled port, the user id to use for authentication
| password   | string | If connecting to a SASL enabled port, the password to use for authentication 
| trustStore | string | If connecting to a TLS secured port, the directory containing the certificates representing the trust chain for the connection. This is usually just the CACert used to sign the server's certificate

### Input:

| Name       | Type   | Description
|:---        | :---   | :---  
| message    | string | The message to send 

### Output:

| Name         | Type     | Description
|:---          | :---     | :---   
| partition    | int32    | Documents the partition that the message was placed on
| offSet       | int64    | Documents the offset for the message

## Examples

The below example sends `Hello From Flogo` to a Kafka Broker running on localhost:

```json
{
  "id": "publish_kafka_message",
  "name": "Publish Message to Kafka",
  "activity": {
    "ref": "github.com/project-flogo/contrib/activity/kafka",
    "settings": {
      "brokerUrls" : "localhost:9092",
      "topic"      : "syslog"
    },
    "input": {
      "message"    : "Hello From Flogo"
    }
  }
}
```
 
## Development

### Testing

To run tests first set up the Kafka broker using the docker-compose file given below:

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
