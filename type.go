package serializer

import (
	"io"
	"reflect"
)

func NewType(v interface{}) Type {
	t := reflect.TypeOf(v)

	return Type{
		newTypeID(t),
		newConstructor(t),
	}
}

type Type struct {
	ID          TypeID
	Constructor Constructor
}

type TypeID string

func NewTypeID(v interface{}) TypeID {
	return newTypeID(reflect.TypeOf(v))
}

func newTypeID(t reflect.Type) TypeID {
	delimiter := "."
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		delimiter = ".*"
	}
	return TypeID(t.PkgPath() + delimiter + t.Name())
}

type Constructor interface {
	New(r io.Reader, decoder Decoder) (interface{}, error)
}

type constructor struct {
	isPtr bool
	typ   reflect.Type
}

func NewConstructor(v interface{}) Constructor {
	return newConstructor(reflect.TypeOf(v))
}

func newConstructor(t reflect.Type) constructor {
	isPtr := false
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		isPtr = true
	}

	return constructor{
		isPtr,
		t,
	}
}

func (c constructor) newPointer() interface{} {
	return reflect.New(c.typ).Interface()
}

func (c constructor) restore(v interface{}) interface{} {
	if c.isPtr {
		return v
	}
	return reflect.ValueOf(v).Elem().Interface()
}

func (c constructor) New(r io.Reader, decoder Decoder) (interface{}, error) {
	v := c.newPointer()
	if err := decoder.Decode(r, v); err != nil {
		return nil, err
	}
	return c.restore(v), nil
}
