# serializer 

[![CircleCI](https://circleci.com/gh/morikuni/serializer/tree/master.svg?style=shield)](https://circleci.com/gh/morikuni/serializer/tree/master)
[![GoDoc](https://godoc.org/github.com/morikuni/serializer?status.svg)](https://godoc.org/github.com/morikuni/serializer)
[![Go Report Card](https://goreportcard.com/badge/github.com/morikuni/serializer)](https://goreportcard.com/report/github.com/morikuni/serializer)
[![codecov](https://codecov.io/gh/morikuni/serializer/branch/master/graph/badge.svg)](https://codecov.io/gh/morikuni/serializer)

serializer serializes object into bytes, and deserializes object from bytes.

## Supported format

- JSON
- Protocol Buffer (require object implementing `proto.Message`)
- Custom format (implement `Marshaler`/`Encoder`)

## Example

```go
package main

import (
	"bytes"
	"fmt"

	"github.com/morikuni/serializer"
)

type Object struct {
	ID   int64
	Name string
}

func main() {
	s := serializer.New()

	// Register the object before use it.
	if err := s.Register(Object{}); err != nil {
		panic(err)
	}

	obj := Object{
		12345,
		"foo",
	}

	buf := &bytes.Buffer{}
	// Serialize to io.Writer
	err := s.Serialize(buf, obj)
	if err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
	// {"id":"main.Object","payload":{"ID":12345,"Name":"foo"}}

	// Deserialize from a io.Reader without specifying the actual type.
	x, err := s.Deserialize(buf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", x)
	// main.Object{ID:12345, Name:"foo"}
}
```
