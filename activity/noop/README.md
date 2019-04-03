<!--
title: No-Op
weight: 4615
-->

# No-Op
This activity is a simple No-Op that can be used for testing.

## Installation

### Flogo CLI
```bash
flogo install github.com/project-flogo/contrib/activity/noop
```

## Examples
Configuration of a No-Op activity

```json
{
  "id": "noop",
  "name": "NoOp",
  "activity": {
    "ref": "github.com/project-flogo/contrib/activity/noop"
  }
}
```