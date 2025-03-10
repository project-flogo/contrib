# JSON2XML

Activity for converting JSON object into XML.

For additional tuneups and syntax take a look at [MXJ](https://github.com/clbanning/mxj) Go package.


## Installation

### Flogo CLI

```bash
flogo install github.com/project-flogo/contrib/activity/json2xml
```

## Configuration

### Input

| Name       | Type    | Description                   |
|------------|---------|-------------------------------|
| json       | object  | Input object in JSON format   |
| xmlRootTag | string  | Optional name of XML root tag |


### Output 

| Name | Type | Description    |
|------|------|----------------|
| xml  | byte | Raw XML Output |

Example converting raw XML byte array into string. 
```
string.tostring($activity[JsonToXML].xmlData )
```

## Usage

```json
{
    "id": "JsonToXML",
    "name": "JsonToXML",
    "activity": {
        "ref": "github.com/project-flogo/contrib/activity/json2xml",
        "input": {
            "xmlRootTag": "",
            "json": "{\"hello\":\"world\"}"
        }
    }
}
```

Output:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<hello>world</hello>
```