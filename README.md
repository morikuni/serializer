# serializer 

[![GoDoc](https://godoc.org/github.com/morikuni/serializer?status.svg)](https://godoc.org/github.com/morikuni/serializer)

serializer serializes a Go object into a bytes.

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
	s := serializer.NewSerializer(
		serializer.NewJSONMarshaler(),
		serializer.NewTextJSONEncoder(),
	)

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

	fmt.Println(buf.String())
	// {"name":"main.Object","payload":"{\"ID\":12345,\"Name\":\"foo\"}"}

	// Deserialize from a io.Reader without specifying the actual type.
	x, err := s.Deserialize(buf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", x)
	// main.Object{ID:12345, Name:"foo"}
}
}
```
