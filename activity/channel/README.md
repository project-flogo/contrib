<!-- 
title: Channel
weight: 4603
-->

# Channel
This activity allows you to put a data on a named channel in the flogo engine.  Channels are
essentially an internal communication channel in the engine.


## Installation

### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/activity/channel
```

## Configuration

### Input:
| Name    | Type   | Description
|:---     | :---   | :---    
| channel | string | The name of channel to use - **REQUIRED**
| data    | any    | The data to put on the channel

