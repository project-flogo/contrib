<!--
title: JSON
weight: 4601
-->

# JSON Functions
This function package exposes common json functions.

## exists()
Check if the key/JSONPath is present in the json object. Refer https://github.com/oliveagle/jsonpath for expression format.
### Input Args

| Arg      | Type   | Description                  |
|:---------|:-------|:-----------------------------|
| object   | any    | The json object              |
| key      | string | The key or JSONPath to check |

### Output

| Arg       | Type   | Description                                         |
|:----------|:-------|:----------------------------------------------------|
| returnVal | bool   | True if value is associated with the key / JSONPath |

## get()
Get value of associated key from json object

### Input Args

| Arg      | Type   | Description                        |
|:---------|:-------|:-----------------------------------|
| object   | any    | The json object                    |
| key      | string | The key for which the value to get |

### Output

| Arg       | Type   | Description                       |
|:----------|:-------|:----------------------------------|
| returnVal | any    | The value associated with the key |


## length()
Get the number of top level elements in json object or array

### Input Args

| Arg               | Type   | Description                        |
|:------------------|:-------|:-----------------------------------|
| jsonObjectOrArray | any    | The json object or array           |

### Output

| Arg           | Type   | Description                            |
|:--------------|:-------|:---------------------------------------|
| lengthInteger | int    | The top level items in the json object |


## numbersToString()
Convert every number type to string in a json object or array

### Input Args

| Arg                 | Type   | Description                        |
|:--------------------|:-------|:-----------------------------------|
| jsonObjectOrArray   | any    | The json object or array           |

### Output

| Arg               | Type   | Description                                              |
|:------------------|:-------|:---------------------------------------------------------|
| jsonObjectOrArray | any    | The json object or array with numbers encoded as strings |


## objKeys()
Get the list of all top level keys of json object

### Input Args

| Arg      | Type   | Description                        |
|:---------|:-------|:-----------------------------------|
| object   | any    | The json object                    |

### Output

| Arg            | Type   | Description                |
|:---------------|:-------|:---------------------------|
| arrayOfStrings | array  | The list of top level keys |


## objValues()
Get the list of all top level values of json object

### Input Args

| Arg      | Type   | Description                        |
|:---------|:-------|:-----------------------------------|
| object   | any    | The json object                    |

### Output

| Arg           | Type   | Description                                     |
|:--------------|:-------|:------------------------------------------------|
| arrayOfValues | array  | The list of all top level values in json object |


## path()
Apply a JSON path to an object.

### Input Args

| Arg       | Type   | Description                          |
|:----------|:-------|:-------------------------------------|    
| path      | string | The JSON path                        |
| object    | any    | The object to apply the JSON path to |

### Output

| Arg       | Type   | Description                 |
|:----------|:-------|:----------------------------|    
| returnVal | any    | The result of the JSON path |


## set()
Set the value of existing key or add new key and set it's value in a json object

### Input Args

| Arg      | Type   | Description                        |
|:---------|:-------|:-----------------------------------|
| object   | any    | The json object                    |
| key      | string | The key for value                  |
| value    | any    | The value for key                  |

### Output

| Arg        | Type   | Description             |
|:-----------|:-------|:------------------------|
| jsonObject | any    | The updated json object |


## jq()
Process the JSON inputs based on a provided query and return the filtered output

### Input Args

| Arg      | Type   | Description                        |
|:---------|:-------|:-----------------------------------|
| object   | any    | The json object                    |
| query    | string | Query to perform                   |

### Output

| Arg        | Type   | Description             |
|:-----------|:-------|:------------------------|
| output     | array  | result of the jq filter |