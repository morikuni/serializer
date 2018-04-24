package serializer

import (
	"fmt"
	"path"
	"reflect"
)

// TypeNameResolver resolve the type name of v.
type TypeNameResolver interface {
	ResolveName(v interface{}) string
}

// NewTypeNameResolver resolve a name as a package path.
func NewTypeNameResolver() TypeNameResolver {
	return resolver{}
}

type resolver struct{}

func (resolver) ResolveName(v interface{}) string {
	return typeNameOf(v)
}

func typeNameOf(v interface{}) string {
	t := reflect.TypeOf(v)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	dir := path.Dir(t.PkgPath())

	return path.Join(dir, fmt.Sprintf("%T", v))
}
