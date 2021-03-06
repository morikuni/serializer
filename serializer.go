// Package serializer serializes object into bytes,
// and deserializes object from bytes.
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
		NewJSONEncoder(),
		NewTypeNameResolver(),
	}

	for _, opt := range opts {
		opt(&s)
	}

	return s
}

// Serializer serialize and deserialize a object.
type Serializer struct {
	typeMap  map[string]reflect.Type
	encoder  Encoder
	resolver TypeNameResolver
}

// Register registers the types into serializer.
func (s Serializer) Register(types ...interface{}) error {
	for _, typ := range types {
		t := reflect.TypeOf(typ)
		name, aliases := s.resolver.ResolveName(typ)

		if err := s.register(name, t); err != nil {
			return err
		}

		for _, a := range aliases {
			if err := s.register(a, t); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s Serializer) register(name string, t reflect.Type) error {
	if old, ok := s.typeMap[name]; ok {
		return DuplicatedNameError{name, old, t}
	}
	s.typeMap[name] = t
	return nil
}

// Serialize serialized a object into w.
func (s Serializer) Serialize(w io.Writer, v interface{}) error {
	name, _ := s.resolver.ResolveName(v)
	if _, ok := s.typeMap[name]; !ok {
		return UnknownTypeError{name}
	}

	return s.encoder.Encode(w, name, v)
}

// Deserialize deserializes a object from r.
func (s Serializer) Deserialize(r io.Reader) (interface{}, error) {
	name, payload, err := s.encoder.Decode(r)
	if err != nil {
		return nil, err
	}

	t, ok := s.typeMap[name]
	if !ok {
		return nil, UnknownTypeError{name}
	}

	isPtr := false
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		isPtr = true
	}

	i := reflect.New(t).Interface()

	if err := s.encoder.Unmarshal(payload, i); err != nil {
		return nil, err
	}

	if !isPtr {
		i = reflect.ValueOf(i).Elem().Interface()
	}

	return i, nil
}
