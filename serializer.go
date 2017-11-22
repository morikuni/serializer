package serializer

import (
	"bytes"
	"io"
)

type Serializer interface {
	Register(types ...interface{})
	Serialize(w io.Writer, v interface{}) error
	Deserialize(r io.Reader) (interface{}, error)
	SerializeByte(v interface{}) ([]byte, error)
	DeserializeByte(data []byte) (interface{}, error)
}

func NewSerializer(options ...Option) Serializer {
	s := serializer{
		NewTypeRegistry(),
		NewJSONCodec(),
	}
	for _, o := range options {
		o(&s)
	}
	return s
}

type serializer struct {
	registry TypeRegistry
	codec    Codec
}

func (s serializer) Register(types ...interface{}) {
	for _, t := range types {
		s.registry.Register(NewType(t))
	}
}

func (s serializer) Serialize(w io.Writer, v interface{}) error {
	id := NewTypeID(v)
	if _, ok := s.registry.Find(id); !ok {
		return UnknownTypeError{id}
	}

	buf := &bytes.Buffer{}
	if err := s.codec.Encode(buf, v); err != nil {
		return err
	}

	d := Data{
		id,
		buf.Bytes(),
	}
	return s.codec.Encode(w, &d)
}

func (s serializer) Deserialize(r io.Reader) (interface{}, error) {
	var d Data
	if err := s.codec.Decode(r, &d); err != nil {
		return nil, err
	}

	ctor, ok := s.registry.Find(d.ID)
	if !ok {
		return nil, UnknownTypeError{d.ID}
	}

	buf := bytes.NewReader(d.Payload)
	return ctor.New(buf, s.codec)
}

func (s serializer) SerializeByte(v interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := s.Serialize(buf, v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s serializer) DeserializeByte(data []byte) (interface{}, error) {
	buf := bytes.NewReader(data)
	return s.Deserialize(buf)
}
