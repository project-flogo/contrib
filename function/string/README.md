<!--
title: String
weight: 4601
-->

# String Functions
This function package exposes common string related functions.

## concat()
Concatenate a set of strings.

### Input Args

| Arg     | Type   | Description
|:---      | :---   | :---    
| str | string | Strings to concatinate

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | string | A concatinated string.

## equals()
Check if two strings are equal

### Input Args

| Arg     | Type   | Description
|:---      | :---   | :---    
| str1 | string | String to be compared 
| str2 | string | String to be compared

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | bool | True if the strings are equal, otherwise false.

## equalsIgnoreCase()
Check if two strings are equal, ignoring case.

### Input Args

| Arg     | Type   | Description
|:---      | :---   | :---    
| str1 | string | String to be compared 
| str2 | string | String to be compared

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | bool | True if the strings are equal, otherwise false.

## contains()
Check if str2 is within str1.

### Input Args

| Arg     | Type   | Description
|:---      | :---   | :---    
| str1 | string | Source string
| str2 | string | String to find in str1

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | bool | True if the str2 is found within str1

## float()
Convert str1 to a foat64.

### Input Args

| Arg     | Type   | Description
|:---      | :---   | :---    
| str1 | string | Source string

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | float64 | The float value of str1

## integer()
Convert str1 to a int.

### Input Args

| Arg     | Type   | Description
|:---      | :---   | :---    
| str1 | string | Source string

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | int | The int value of str1

## len()
Get the length of str1

### Input Args

| Arg     | Type   | Description
|:---      | :---   | :---    
| str1 | string | Source string

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | int | The length of str1

## substring()
Get a substring from str1

### Input Args

| Arg     | Type   | Description
|:---      | :---   | :---    
| str1 | string | Source string
| start | string | The starting string/char
| start | string | The ending string/char

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | string | The substring