exregexp-go
===========

[![Go Reference](https://pkg.go.dev/badge/github.com/hymkor/exregexp-go.svg)](https://pkg.go.dev/github.com/hymkor/exregexp-go)
[![Go Test](https://github.com/hymkor/exregexp-go/actions/workflows/go.yml/badge.svg)](https://github.com/hymkor/exregexp-go/actions/workflows/go.yml)

`exregexp-go` is a Go package that extends the standard `regexp` library by providing utilities for flexible string replacement using regular expression submatches.

This package introduces the `ReplaceAllStringSubmatchFunc` function, which allows you to leverage capturing groups for custom replacement logic.

Features
--------

- Flexible string replacement using capturing groups
- Simple and intuitive API
- Built on top of Go's standard `regexp` library

Installation
------------

```sh
go get -u github.com/hymkor/exregexp-go
```

Usage
-----

The following example demonstrates how to use `exregexp.ReplaceAllStringSubmatchFunc`:

```example.go
package main

import (
    "fmt"
    "regexp"

    "github.com/hymkor/exregexp-go"
)

func main() {
    re := regexp.MustCompile(`\b([a-zA-Z]+)(\d+)\b`)
    input := "example123 test456 hello789"
    output := exregexp.ReplaceAllStringSubmatchFunc(re, input, func(submatches []string) string {
        return fmt.Sprintf("%s(%s)", submatches[1], submatches[2])
    })

    fmt.Println(output)
}
```

```go run example.go|
example(123) test(456) hello(789)
```

Author
------

- [hymkor (HAYAMA Kaoru)](https://github.com/hymkor)

LICENSE
-------

- [MIT LICENSE](./LICENSE)
