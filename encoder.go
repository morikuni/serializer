package serializer

import "io"

type Encoder interface {
	Encode(w io.Writer, d Data) error
	Decode(r io.Reader) (Data, error)
}
