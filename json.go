package serializer

import (
	"encoding/json"
	"io"
)

func NewJSONCodec() Codec {
	return jsonEncoder{}
}

type jsonEncoder struct{}

func (e jsonEncoder) Encode(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func (e jsonEncoder) Decode(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
