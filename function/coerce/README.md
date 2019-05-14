<!--
title: Coerce
weight: 4601
-->

# Data Coerce Functions
This function package enables you to coerce data from one type to another.

Valid data types in Flogo include:

| Type     | Description
|:---      | :---    
| any | Any data type.
| string | A string
| int | A standard int
| int32 | A 32 bit integer
| int64 | A 64 bit integer
| float32 | A 32 bit float
| float64 | A 64 bit float
| bool | A boolean (true | false)
| object | A flogo object. Essentially a JOSN object
| bytes | A byte data type
| params | Parameters data type
| array | An array type
| map | A map

## toType()
Used to convert a value to a specified type.

### Input Args

| Arg     | Type   | Description
|:---      | :---   | :---    
| value | any | The value to be coerced.
| type | string | The data type to coerce to.

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | any | The coerced value

## toString()
Convert the spcified value to a string.

### Input Args

| Arg     | Type   | Description
|:---      | :---   | :---    
| value | any | The value to be coerced.

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | string | The coerced string value

## toInt()
Convert the specified value to an int.

### Input Args

| Arg     | Type   | Description
|:---      | :---   | :---    
| value | any | The value to be coerced.

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | int | The coerced int value

## toInt32()
Convert the specified value to an int32.

### Input Args

| Arg     | Type   | Description
|:---      | :---   | :---    
| value | any | The value to be coerced.

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | int32 | The coerced int32 value

## toInt64()
Convert the specified value to a int64.

### Input Args

| Arg     | Type   | Description
|:---      | :---   | :---    
| value | any | The value to be coerced.

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | int64 | The coerced int64 value

## toFloat32()
Convert the specified value to a float32.

### Input Args

| Arg     | Type   | Description
|:---      | :---   | :---    
| value | any | The value to be coerced.

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | float32 | The coerced int32 value

## toFloat64()
Convert the specified value to a float64.

### Input Args

| Arg     | Type   | Description
|:---      | :---   | :---    
| value | any | The value to be coerced.

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | float64 | The coerced float64 value

## toBool()
Convert the specified value to a boolean.

### Input Args

| Arg     | Type   | Description
|:---      | :---   | :---    
| value | any | The value to be coerced.

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | bool | The coerced bool value

## toBytes()
Convert the specified value to bytes.

### Input Args

| Arg     | Type   | Description
|:---      | :---   | :---    
| value | any | The value to be coerced.

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | byte | The coerced byte value

## toParams()
Convert the specified value to params type.

### Input Args

| Arg     | Type   | Description
|:---      | :---   | :---    
| value | any | The value to be coerced.

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | params | The coerced params value

## toObject()
Convert the specified value to an object type.

### Input Args

| Arg     | Type   | Description
|:---      | :---   | :---    
| value | any | The value to be coerced.

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | object | The coerced object

## toArray()
Convert the specified value to an array.

### Input Args

| Arg     | Type   | Description
|:---      | :---   | :---    
| value | any | The value to be coerced.

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | array | The coerced array