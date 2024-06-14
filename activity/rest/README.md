# REST Activity
This activity allows you to invoke a REST service. 
It will pass through the uri and reads the domain:port from an application configuration if configured via the path settings, e.g.:
`{schemadomainnameport}/` where `schemadomainnameport` is the application property value.
The setting  `` needs to be set to yes to function.

# Installation
flogo install github.com/project-flogo/contrib/activity/rest

## TIBCO sandbox manual installation
```bash
zip the 'rest' directory (including the directory 'rest' )
Go to https://eu.integration.cloud.tibco.com/envtools/flogo_extensions and upload the zip
ensure to increase the version nr in the description.json before uploading to distingues between versions
```

# Configuration
## Settings:
| Name       | Type    | Description                                                                                                                                                                                                  |
|------------|---------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| method     | string  | The HTTP method to invoke (Allowed values are GET, POST, PUT, DELETE, PATCH, TRIGGER) - **REQUIRED** <br/>When TRIGGER is selected, the actual method is taken from the incoming trigger                     |
| useEnvProp | boolean | define if you want to replace the schema://hostname:port part with the environment variable                                                                                                                  |
| uri        | string  | The URI of the service to invoke - **REQUIRED** <br> When the uri start with an { it is seen as the start of a applciation property. So {ap} where ap is the applciation property defined to replace runtime |
| headers    | params  | The HTTP header parameters                                                                                                                                                                                   |
| proxy      | string  | The address of the proxy server to be used                                                                                                                                                                   |
| timeout    | int     | The request timeout in seconds                                                                                                                                                                               |
| sslConfig  | object  | SSL configuration                                                                                                                                                                                            |


## *sslConfig* Object: 
| Property      | Type   | Description                                             |
|---------------|--------|---------------------------------------------------------|
| skipVerify    | bool   | Skip SSL validation, defaults to true                   |
| useSystemCert | bool   | Use the systems root certificate file, defaults to true |
| caFile        | string | The path to PEM encoded root certificates file          |
| certFile      | string | The path to PEM encoded client certificate              |
| keyFile       | string | The path to PEM encoded client key                      |

*Note: used if URI is https*
## Input:
| Name        | Type   | Description                                                                |
|-------------|--------|----------------------------------------------------------------------------|
| pathParams  | params | The path parameters (e.g., 'id' in http://.../pet/:id/name )               |
| queryParams | params | The query parameters (e.g., 'id' in http://.../pet?id=someValue )          |
| headers     | params | The HTTP header parameters                                                 |
| method      | string | the method coming in from the trigger; Only used if 'TRIGGER'  is selected |
| content     | any    | The message content to send. This is only used in POST, PUT, and PATCH     |
| envPropUri  | string | Connect the availalbe environment variable into this activity              |


## Output:
| Name    | Type   | Description                              |
|---------|--------|------------------------------------------|
| status  | int    | The HTTP status code                     |
| data    | any    | The HTTP response data                   |
| headers | params | The HTTP response headers                |
| cookies | array  | The response cookies (from 'Set-Cookie') |
