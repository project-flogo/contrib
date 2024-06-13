# REST Trigger
This trigger provides your flogo application the ability to start an action via REST over HTTP
If the option `isPassThroughUri` is set, a pathParam is automatically create with the rest of the path after the give settings path of the handler
## Installation

```bash
flogo install github.com/project-flogo/contrib/trigger/rest
```

## TIBCO sandbox - Manual loading
```bash
zip the 'rest' directory (including the directory 'rest' )
Go to https://eu.integration.cloud.tibco.com/envtools/flogo_extensions and upload the zip
ensure to increase the version nr in the description.json before uploading to distingues between versions
```

## Configuration

### Settings:
| Name      | Type   | Description                                |
|:----------|:-------|:-------------------------------------------|
| port      | int    | The port to listen on - **REQUIRED**       |
| enableTLS | bool   | Enable TLS on the server                   |
| certFile  | string | The path to PEM encoded server certificate |
| keyFile   | string | The path to PEM encoded server key         |

### Handler Settings:
| Name       | Type    | Description                                                        |
|:-----------|:--------|:-------------------------------------------------------------------|
| method     | string  | The HTTP method (ie. GET,POST,PUT,PATCH or DELETE) - **REQUIRED**  |
| path       | string  | The resource path - **REQUIRED**                                   |


### Output:
| Name        | Type   | Description                                                       |
|-------------|--------|-------------------------------------------------------------------|
| pathParams  | params | The path parameters (e.g., 'id' in http://.../pet/:id/name )      |
| queryParams | params | The query parameters (e.g., 'id' in http://.../pet?id=someValue ) |
| headers     | params | The HTTP header parameters                                        |
| method      | string | The HTTP method used for the request                              |
| content     | any    | The content of the request                                        |

### Reply:
| Name    | Type   | Description                                                  |
|---------|--------|--------------------------------------------------------------|
| code    | int    | The http code to reply with                                  |
| data    | any    | The data to reply with                                       |
| headers | params | The HTTP response headers                                    |
| cookies | params | The HTTP response cookies to set (uses 'Set-Cookie' headers) |
