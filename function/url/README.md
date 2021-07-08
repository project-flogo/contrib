<!--
title: URL
weight: 4601
-->

# URL Functions
This function package exposes common URL functions.

## encode()
Returns the URL encoded form of input string

### Input Args

| Arg          | Type   | Description    |
|:-------------|:-------|:---------------|
| rawURLString | string | Raw URL string |

### Output

| Arg       | Type   | Description            |
|:----------|:-------|:-----------------------|
| returnVal | string | The url encoded string |


## escapedPath()
Returns the escaped path part of the URL, stripping everything except the PATH after hostname.

### Input Args

| Arg          | Type   | Description    |
|:-------------|:-------|:---------------|
| rawURLString | string | Raw URL string |

### Output

| Arg       | Type   | Description                               |
|:----------|:-------|:------------------------------------------|
| returnVal | string | The escaped PATH part of the rawURLString |


## hostname()
Returns hostname for the URL, stripping any valid port number if present. If input is enclosed in square brackets, as literal IPv6 addresses are, the square brackets are removed from the output.

### Input Args

| Arg          | Type   | Description    |
|:-------------|:-------|:---------------|
| rawURLString | string | Raw URL string |

### Output

| Arg       | Type   | Description                      |
|:----------|:-------|:---------------------------------|
| returnVal | string | The hostname of the rawURLString |


## path()
Returns the path part of URL

### Input Args

| Arg          | Type    | Description    |
|:-------------|:--------|:---------------|
| rawURLString | string  | Raw URL string |

### Output

| Arg       | Type   | Description          |
|:----------|:-------|:---------------------|
| returnVal | any    | The path part of URL |


## pathEscape()
Returns the escaped string so it can be safely placed inside a URL path segment, replacing special characters (including /) with %XX sequences as needed.

### Input Args

| Arg        | Type   | Description     |
|:-----------|:-------|:----------------|
| pathString | string | The path string |

### Output

| Arg       | Type   | Description             |
|:----------|:-------|:------------------------|
| returnVal | string | The escaped PATH string |


## port()
Returns the port part of URL, without the leading colon. If URL doesn't contain a valid numeric port, port returns an empty string.

### Input Args

| Arg          | Type   | Description    |
|:-------------|:-------|:---------------|
| rawURLString | string | Raw URL string |

### Output

| Arg       | Type   | Description          |
|:----------|:-------|:---------------------|
| returnVal | string | The port part of URL |


## query()
Returns the encoded query string if second parameter is true else returns an object with key value pair of query and value

### Input Args

| Arg          | Type    | Description                                          |
|:-------------|:--------|:-----------------------------------------------------|
| rawURLString | string  | Raw URL string                                       |
| encode       | boolean | Set true to encode query string, false to get object |

### Output

| Arg       | Type   | Description                        |
|:----------|:-------|:-----------------------------------|
| returnVal | any    | The encoded query string or object |


## queryEscape()
Encodes the input string so it can be safely placed inside a URL query. Please note, this does not create the full query string.

### Input Args

| Arg        | Type   | Description     |
|:-----------|:-------|:----------------|
| queryValue | string | URL query value |

### Output

| Arg       | Type   | Description                    |
|:----------|:-------|:-------------------------------|
| returnVal | string | The escaped value of the input |


## scheme()
Returns the URL scheme

### Input Args

| Arg          | Type   | Description    |
|:-------------|:-------|:---------------|
| rawURLString | string | Raw URL string |

### Output

| Arg       | Type   | Description    |
|:----------|:-------|:---------------|
| returnVal | string | The URL scheme |
