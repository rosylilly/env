# github.com/rosylilly/env

env provides utility methods for environment variables.

## Example

```go
package main

import (
  "log"

  "github.com/rosylilly/env"
)

type JSONValue struct {
  A string `json:"a"`
  B int    `json:"b"`
}

func main() {
  strVal := env.String("FOO", "default value")
  log.Println(strVal)

  intVal := env.Int("BAR", 100)
  log.Println(intVal)

  jsonVal := env.JSON("BAZ", JSONValue{})
  log.Println(jsonVal)
}
```
