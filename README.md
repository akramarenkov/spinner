# Spinner

[![Go Reference](https://pkg.go.dev/badge/github.com/akramarenkov/spinner.svg)](https://pkg.go.dev/github.com/akramarenkov/spinner)
[![Go Report Card](https://goreportcard.com/badge/github.com/akramarenkov/spinner)](https://goreportcard.com/report/github.com/akramarenkov/spinner)
[![Coverage Status](https://coveralls.io/repos/github/akramarenkov/spinner/badge.svg)](https://coveralls.io/github/akramarenkov/spinner)

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
    spn := spinner.New(-1, 2)
    fmt.Println(spn.Actual())

    for range 7 {
        spn.Spin()
        fmt.Println(spn.Actual())
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
