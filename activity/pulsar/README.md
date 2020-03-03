
# Apache Pulsar Activity

This activity publishes messages on a topic in a Apache Pulsar.

### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/activity/pulsar
```

## Configuration

### Settings: 
| Name       | Type   | Description
|:---        | :---   | :---   
| url        | string | The url of Apache Pulsar to connnect to - ***REQUIRED***
| topic      | string | The Pulsar topic on which to place the message - ***REQUIRED***
| athenzauth | string | If using Athen Authentication, the auth-params provided
| certFile   | string | If using authenitication using TLS, the path of certificate file.
| keyFile    | string | If using authenitication using TLS, the path of private Key file
### Input:

| Name       | Type   | Description
|:---        | :---   | :---  
| payload    | string | The message to send 

