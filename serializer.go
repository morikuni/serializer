// Package serializer serializes a struct into bytes.
package serializer

import (
	"io"
	"reflect"
)

// New returns a Serializer with given marshaler
// and encoder.
func New(opts ...Option) Serializer {
	s := Serializer{
		make(map[string]reflect.Type),
		NewJSONMarshaler(),
		NewProtobufEncoder(),
		NewTypeNameResolver(),
	}

	for _, opt := range opts {
		opt(&s)
	}

	return s
}

// Serializer serialize and deserialize a object.
type Serializer struct {
	typeMap   map[string]reflect.Type
	marshaler Marshaler
	encoder   Encoder
	resolver  TypeNameResolver
}

// Register registers the types into serializer.
func (s Serializer) Register(types ...interface{}) {
	for _, t := range types {
		typ := reflect.TypeOf(t)
		name, aliases := s.resolver.ResolveName(t)
		s.typeMap[name] = typ
		for _, a := range aliases {
			s.typeMap[a] = typ
		}
	}
}

// Serialize serialized a object into w.
func (s Serializer) Serialize(w io.Writer, v interface{}) error {
	name, _ := s.resolver.ResolveName(v)
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
