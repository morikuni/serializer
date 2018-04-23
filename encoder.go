package serializer

import "io"

// Encoder encodes and decodes a Data.
type Encoder interface {
	Encode(w io.Writer, d Data) error
	Decode(r io.Reader) (Data, error)
}
