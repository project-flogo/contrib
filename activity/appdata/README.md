<!--
title: AppData
weight: 4616
-->

# AppData
This activity allows you to set and get global App attributes.

## Installation

### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/activity/appdata
```

## Configuration

### Settings:
| Name | Type   | Description
|:---  | :---   | :---    
| name | string | The name of the shared attribute - **REQUIRED**         
| op   | string | The operation (get or set), 'get' is the default
| type | string | The data type of the shared value, default is 'any'

### Input:
| Name  | Type   | Description
|:---   | :---   | :---    
| value | object |  The value of the shared attribute


### Output:
| Name  | Type   | Description
|:---   | :---   | :---    
| value | object |  The value of the shared attribute


## Examples

### Get
Get the value of the 'myAttr' attribute:

```json
{
  "id": "get_app_attr",
  "name": "Get App Attr",
  "activity": {
    "ref": "github.com/project-flogo/contrib/activity/appdata",
    "settings": {
      "attribute": "myAttr",
      "operation": "get"
    }
  }
}
```

### Set
Update the value of the 'myAttr' attribute to _bar_:

```json
{
  "id": "set_app_attr",
  "name": "Set App Attr",
  "activity": {
    "ref": "github.com/project-flogo/contrib/activity/appdata",
    "settings": {
      "attribute": "myAttr",
      "operation": "set"
    },
    "input": {
      "value": "bar"
    }  
  }
}
```
