package serializer

import (
	"encoding/json"
	"io"
)

func NewJSONMarshaler() Marshaler {
	return jsonMarshaler{}
}

type jsonMarshaler struct{}

var _ interface {
	Marshaler
} = jsonMarshaler{}

func (jsonMarshaler) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (jsonMarshaler) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func NewJSONEncoder() Encoder {
	return jsonEncoder{}
}

type jsonEncoder struct{}

func (jsonEncoder) Encode(w io.Writer, d Data) error {
	return json.NewEncoder(w).Encode(d)
}

func (jsonEncoder) Decode(r io.Reader) (Data, error) {
	var d Data
	if err := json.NewDecoder(r).Decode(&d); err != nil {
		return Data{}, err
	}
	return d, nil
}
