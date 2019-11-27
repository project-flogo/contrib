<!--
title: HTTP
weight: 4601
-->

# HTTP Functions
This function package exposes http related related utility functions.

## reqCookieToParams()
Converts the 'Cookie' request header value string to Params object.

### Example
```http.reqCookieToParams($.headers["Cookie"]) => { "cookie1":"value1", "cookie2":"value2" }```

### Input

| Arg     | Type   | Description
|:---      | :---   | :---    
| cookies | string | The string from the 'Cookie' request header

### Output

| Type   | Description
| :---   | :---    
| params | A string map of cookie name to value

## reqCookieFromParams()
Converts a Params object to a string suitable for setting 'Cookie' request header value.

#### Example

```http.reqCookieFromParams($flow.myCookies) => "cookie1=value1;cookie2=value2"```

### Input

| Arg     | Type   | Description
|:---      | :---   | :---    
| cookies | params | The params representation of the cookies from the 'Cookie' request header

### Output

| Type   | Description
| :---   | :---    
| string | A the string representation of the request cookies

## resCookieToObject()
Converts a cookie from a 'Set-Cookie' response header value to an object.

### Example

```http.resCookieToObject($activity[rest].cookies[0]) => { "Name":"cookie1", "Value":"value1", "Secure":true }```

### Input

| Arg     | Type   | Description
|:---    | :---   | :---    
| cookie | string | The string cookie from a 'Set-Cookie' response header

### Output

| Type   | Description
| :---   | :---    
| object | An object representation of a response cookie

## resCookieFromObject()
Converts an Object to a string suitable for setting 'Set-Cookie' response header value.

### Example

```http.resCookieFromObject($flow.myCookie) => "cookie1=value1;Secure"```

### Input

| Arg     | Type   | Description
|:---    | :---   | :---    
| cookie | any | The object representation of a cookie from a 'Set-Cookie' response header, can also be of type map or params

### Output

| Type   | Description
| :---   | :---    
| object | An object representation of a response cookie

## resCookiesToObjectMap()
Converts an array of response cookies to a map of cookie objects.

### Example

```http.resCookiesToObjectMap($activity[rest].cookies) => { "cookie1":{ "Name":"cookie1", "Value":"value1", "Secure":true }}```

### Input

| Arg     | Type   | Description
|:---    | :---   | :---    
| cookies | array | An array of response cookies from the 'Set-Cookie' response headers

### Output

| Type   | Description
| :---   | :---    
| map | A map of cookie objects

## resCookiesFromObjectMap()
Converts a map of cookie objects to an array of cookie strings suitable for setting 'Set-Cookie' response header values.

### Example

```http.resCookiesFromObjectMap($flow.myCookies) => ["cookie1=value1;Secure"] }```

### Input

| Arg     | Type   | Description
|:---    | :---   | :---    
| cookies | map | A map of name to cookie object

### Output

| Type   | Description
| :---   | :---    
| array | An array of cookie strings for suitable for setting 'Set-Cookie' response headers.