package serializer

import (
	"encoding/json"
	"io"
)

// Encoder encodes and decodes a Data.
type Encoder interface {
	Encode(w io.Writer, name string, v interface{}) error
	Decode(r io.Reader) (name string, payload []byte, err error)
	Unmarshal(payload []byte, v interface{}) error
}

// NewJSONEncoder returns a encoder that encodes a
// object to json.
func NewJSONEncoder() Encoder {
	return jsonEncoder{}
}

type jsonEncoder struct{}

func (jsonEncoder) Encode(w io.Writer, name string, v interface{}) error {
	type encodingJSON struct {
		Name    string      `json:"name"`
		Payload interface{} `json:"payload"`
	}

	d := encodingJSON{name, v}
	return json.NewEncoder(w).Encode(d)
}

func (jsonEncoder) Decode(r io.Reader) (string, []byte, error) {
	type decodingJSON struct {
		Name    string          `json:"name"`
		Payload json.RawMessage `json:"payload"`
	}

	var d decodingJSON
	if err := json.NewDecoder(r).Decode(&d); err != nil {
		return "", nil, err
	}
	return d.Name, d.Payload, nil
}

func (jsonEncoder) Unmarshal(payload []byte, v interface{}) error {
	return json.Unmarshal(payload, v)
}
