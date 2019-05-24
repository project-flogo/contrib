<!--
title: XML2JSON
weight: 4615
-->

# Log
This activity allows you to convert XML data to JSON object.

## Installation

### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/activity/xml2json
```

## Configuration

### Input:
| Name       | Type   | Description
|:---        | :---   | :---    
| xmlData    | string | The XML data to be converted

### Output:
| Name       | Type   | Description
|:---        | :---   | :---    
| jsonObject | object | The JSON object

## Examples
The below example logs a message 'test message':

```json
{
  "id": "xml2json_activity",
  "name": "XML2JSON Activity",
  "activity": {
    "ref": "github.com/project-flogo/contrib/activity/xml2json",
    "input": {
      "xmlData": "<?xml version=\"1.0\" encoding=\"UTF-8\"?><hello>world</hello>"
    }
  }
}
```