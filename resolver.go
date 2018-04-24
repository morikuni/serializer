package serializer

import (
	"fmt"
	"path"
	"reflect"
)

// TypeNameResolver resolve the name of v.
type TypeNameResolver interface {
	// ResolveName has to return the name of v.
	// Return aliases only when the v should be
	// deserialized from other names. (e.g. the
	// actual name of v has renamed by refactoring
	// or something.)
	ResolveName(v interface{}) (name string, aliases []string)
}

// NewTypeNameResolver resolve a name using TypeNameOf.
// The resolver supports AliasNamer interface.
func NewTypeNameResolver() TypeNameResolver {
	return resolver{}
}

// AliasNamer represents the type has alias names.
type AliasNamer interface {
	AliasName() []string
}

type resolver struct{}

func (resolver) ResolveName(v interface{}) (string, []string) {
	var aliases []string
	if an, ok := v.(AliasNamer); ok {
		aliases = an.AliasName()
	}
	return TypeNameOf(v), aliases
}

// TypeNameOf returns a unique name of v.
func TypeNameOf(v interface{}) string {
	t := reflect.TypeOf(v)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	dir := path.Dir(t.PkgPath())

	return path.Join(dir, fmt.Sprintf("%T", v))
}
