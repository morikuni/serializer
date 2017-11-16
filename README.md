# serializer 

[![GoDoc](https://godoc.org/github.com/morikuni/serializer?status.svg)](https://godoc.org/github.com/morikuni/serializer)

serialize go struct into bytes, and deserialize them into go struct.


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
	s := serializer.NewSerializer()

	// Register the struct to serialize and deserialize
	s.Register(Object{})

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

	// Default codec is JSON
	fmt.Println(buf.String())
	// {"id":"main.Object","payload":"eyJJRCI6MTIzNDUsIk5hbWUiOiJmb28ifQo="}

	// Deserialize from io.Reader without specifying the actual type.
	x, err := s.Deserialize(buf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", x)
	// main.Object{ID:12345, Name:"foo"}
}
```
