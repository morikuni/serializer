package serializer

import (
	"io"
)

type Decoder interface {
	Decode(r io.Reader, v interface{}) error
}

type Codec interface {
	Encode(w io.Writer, v interface{}) error
	Decoder
}
