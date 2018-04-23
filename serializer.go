// Package serializer serializes a Go object into a bytes.
package serializer

import (
	"io"
	"reflect"
)

type Serializer interface {
	Register(types ...interface{})
	RegisterByName(name string, t interface{})
	Serialize(w io.Writer, v interface{}) error
	Deserialize(r io.Reader) (interface{}, error)
}

func NewSerializer(marshaler Marshaler, encoder Encoder) Serializer {
	s := serializer{
		make(map[string]reflect.Type),
		marshaler,
		encoder,
	}
	return s
}

type serializer struct {
	typeMap   map[string]reflect.Type
	marshaler Marshaler
	encoder   Encoder
}

func (s serializer) Register(types ...interface{}) {
	for _, t := range types {
		s.RegisterByName(TypeNameOf(t), t)
	}
}

func (s serializer) RegisterByName(name string, t interface{}) {
	s.typeMap[name] = reflect.TypeOf(t)
}

func (s serializer) Serialize(w io.Writer, v interface{}) error {
	name := TypeNameOf(v)
	if _, ok := s.typeMap[name]; !ok {
		return UnknownTypeError{name}
	}

	payload, err := s.marshaler.Marshal(v)
	if err != nil {
		return err
	}

	d := Data{
		name,
		payload,
	}
	return s.encoder.Encode(w, d)
}

func (s serializer) Deserialize(r io.Reader) (interface{}, error) {
	d, err := s.encoder.Decode(r)
	if err != nil {
		return nil, err
	}

	t, ok := s.typeMap[d.Name]
	if !ok {
		return nil, UnknownTypeError{d.Name}
	}

	isPtr := false
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		isPtr = true
	}

	i := reflect.New(t).Interface()

	if err := s.marshaler.Unmarshal(d.Payload, i); err != nil {
		return nil, err
	}

	if !isPtr {
		i = reflect.ValueOf(i).Elem().Interface()
	}

	return i, nil
}
