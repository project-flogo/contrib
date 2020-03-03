
# Apache Pulsar Trigger
This trigger provides your flogo application the ability to listen from the Apache Pulsar.

## Installation

```bash
flogo install github.com/project-flogo/contrib/trigger/pulsar
```

## Configuration

### Settings:
| Name      | Type   | Description
|:---       | :---   | :---       
| url       | string | The url of Apache Pulsar to connnect to - ***REQUIRED***
| athenzauth| string | If using Athen Authentication, the auth-params provided
| certFile  | string | The path to PEM encoded server certificate
| keyFile   | string | The path to PEM encoded server key


### Handler Settings:
| Name         | Type   | Description
|:---          | :---   | :---          
| topic        | string | The Pulsar topic from which to get the message - ***REQUIRED***
| subscription | string | The subscription name - **REQUIRED**

### Output:
| Name        | Type   | Description
|:---         | :---   | :---        
| message     | string | The message from the Pulsar.

