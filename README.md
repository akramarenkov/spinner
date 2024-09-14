# Spinner

[![Go Reference](https://pkg.go.dev/badge/github.com/akramarenkov/spinner.svg)](https://pkg.go.dev/github.com/akramarenkov/spinner)
[![Go Report Card](https://goreportcard.com/badge/github.com/akramarenkov/spinner)](https://goreportcard.com/report/github.com/akramarenkov/spinner)
[![codecov](https://codecov.io/gh/akramarenkov/spinner/branch/master/graph/badge.svg?token=6FLHGpPBW2)](https://codecov.io/gh/akramarenkov/spinner)

## Purpose

Library used to infinitely iterate over integer values ​​in a given range

## Usage

Example:

```go
package main

import (
    "fmt"

    "github.com/akramarenkov/spinner"
)

func main() {
    spinner := spinner.New(-1, 2)
    fmt.Println(spinner.Actual())

    for range 7 {
        spinner.Spin()
        fmt.Println(spinner.Actual())
    }

    // Output:
    // -1
    // 0
    // 1
    // 2
    // -1
    // 0
    // 1
    // 2
}
```
