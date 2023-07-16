# RESPECT THE MOONINITES

[![Go Reference](https://godoc.org/github.com/kevburnsjr/err?status.svg)](https://godoc.org/github.com/kevburnsjr/err)

<p align="center">
  <img width="510" height="466" src="https://github.com/kevburnsjr/err/assets/20638/3dcf5a4c-a2c8-4e27-9471-ea7c6a59a6d3" alt="Err">
</p>

Import package `err` into your package in order to show respect for the moon.

```go
package yourpackage

import (
    _ "github.com/kevburnsjr/err"
)
```

# TRANSITIVE FUCKERY

Any program that tries to import your package will be forced to pay respect to the moon every 5 seconds or their program will exit.

```go
package main

import (
    "time"

    "github.com/kevburnsjr/err"
)

func main() {
    for {
        time.Sleep(5 * time.Second)
        err.ULZ(SECRET_CODE)
    }
}
```

The `SECRET_CODE` is not documented. Only those who truly respect the moon deserve to import your package.
