<!--
title: MATH
weight: 4601
-->

# Math Functions
This function package exposes common math functions.

## ceil()
Ceil returns the least integer value greater than or equal to input

### Input Args

| Arg         | Type   | Description      |
|:------------|:-------|:-----------------|
| inputNumber | number | The input number |

### Output

| Arg       | Type   | Description                        |
|:----------|:-------|:-----------------------------------|
| returnVal | number | The ceil value of the input number |


## floor()
Floor returns the greatest integer value less than or equal to input

### Input Args

| Arg         | Type   | Description      |
|:------------|:-------|:-----------------|
| inputNumber | number | The input number |

### Output

| Arg       | Type   | Description                         |
|:----------|:-------|:------------------------------------|
| returnVal | number | The floor value of the input number |


## isNaN()
IsNaN reports whether input is an IEEE 754 "not-a-number" value

### Input Args

| Arg   | Type   | Description       |
|:------|:-------|:------------------|
| input | any    | The input to test |

### Output

| Arg       | Type    | Description                                 |
|:----------|:--------|:--------------------------------------------|
| returnVal | boolean | Is true if input is IEEE 754 "not-a-number" |


## mod()
Mod returns the floating-point remainder of x/y. The magnitude of the result is less than y and its sign agrees with that of x

### Input Args

| Arg | Type   | Description                 |
|:----|:-------|:----------------------------|
| x   | number | The dividend or first input |
| y   | number | The divisor or second input |

### Output

| Arg       | Type   | Description                                       |
|:----------|:-------|:--------------------------------------------------|
| returnVal | number | The remainder of the Euclidean division of x by y |


## round()
Round returns the nearest integer, rounding half away from zero

### Input Args

| Arg         | Type   | Description      |
|:------------|:-------|:-----------------|
| inputNumber | number | The input number |

### Output

| Arg       | Type   | Description                                       |
|:----------|:-------|:--------------------------------------------------|
| returnVal | number | The nearest integer, rounding half away from zero |


## roundToEven()
RoundToEven returns the nearest integer, rounding ties to even

### Input Args

| Arg         | Type   | Description      |
|:------------|:-------|:-----------------|
| inputNumber | number | The input number |

### Output

| Arg       | Type   | Description                                |
|:----------|:-------|:-------------------------------------------|
| returnVal | number | The nearest integer, rounding ties to even |


## trunc()
Trunc returns the integer value of input

### Input Args

| Arg | Type   | Description                 |
|:----|:-------|:----------------------------|
| inputNumber | number | The input number |

### Output

| Arg       | Type   | Description                              |
|:----------|:-------|:-----------------------------------------|
| returnVal | number | The truncated integer value of the input |