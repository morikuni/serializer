// Package serializer serializes a struct into bytes.
package serializer

import (
	"io"
	"reflect"
)

// NewSerializer returns a Serializer with given marshaler
// and encoder.
func NewSerializer(marshaler Marshaler, encoder Encoder) Serializer {
	s := Serializer{
		make(map[string]reflect.Type),
		marshaler,
		encoder,
	}
	return s
}

// Serializer serialize and deserialize a object.
type Serializer struct {
	typeMap   map[string]reflect.Type
	marshaler Marshaler
	encoder   Encoder
}

// Register registers the types into serializer.
func (s Serializer) Register(types ...interface{}) {
	for _, t := range types {
		s.RegisterByName(TypeNameOf(t), t)
	}
}

// RegisterByName registers the type with given name.
func (s Serializer) RegisterByName(name string, t interface{}) {
	s.typeMap[name] = reflect.TypeOf(t)
}

// Serialize serialized a object into w.
func (s Serializer) Serialize(w io.Writer, v interface{}) error {
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

// Deserialize deserializes a object from r.
func (s Serializer) Deserialize(r io.Reader) (interface{}, error) {
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
