package serializer

import (
	"encoding/json"
	"io"
)

// NewJSONMarshaler returns a marshaler that marshals an
// object to json.
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

// NewJSONEncoder returns a encoder that encodes a
// data to json.
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

// NewTextJSONEncoder returns a encoder that encodes a
// data to json without encoding payload to base64.
func NewTextJSONEncoder() Encoder {
	return textJSONEncoder{}
}

type textJSONEncoder struct{}

func (textJSONEncoder) Encode(w io.Writer, d Data) error {
	sd := stringData{d.Name, string(d.Payload)}
	return json.NewEncoder(w).Encode(sd)
}

func (textJSONEncoder) Decode(r io.Reader) (Data, error) {
	var sd stringData
	if err := json.NewDecoder(r).Decode(&sd); err != nil {
		return Data{}, err
	}
	return Data{sd.Name, []byte(sd.Payload)}, nil
}

type stringData struct {
	Name    string `json:"name"`
	Payload string `json:"payload"`
}
