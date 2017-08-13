[![godoc](https://godoc.org/github.com/haggen/faye?status.svg)](https://godoc.org/github.com/haggen/faye)
[![Go Report Card](https://goreportcard.com/badge/github.com/haggen/faye)](https://goreportcard.com/report/github.com/haggen/faye)
[![CircleCI](https://circleci.com/gh/haggen/faye/tree/master.svg?style=shield)](https://circleci.com/gh/haggen/faye/tree/master)
[![codecov](https://codecov.io/gh/haggen/faye/branch/master/graph/badge.svg)](https://codecov.io/gh/haggen/faye)


# Faye

> Faye client in Go.

## About

See [Faye](https://faye.jcoglan.com/).

## Example

```go
package main

import "fmt"
import "github.com/haggen/faye"

func main() {
    f, err := faye.New("https://example.org/faye")
    if err != nil {
      panic(err)
    }
    channel := f.Subscribe("/channel")
    for {
      m := <-channel
      fmt.Println(m.Data)
    }
}
```

## Legal

Faye (Go) package is under The MIT License © 2017 Arthur Corenzan.

Faye protocol is under The MIT License © 2009–2017 James Coglan.
