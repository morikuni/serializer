package serializer

import (
	"reflect"
)

func TypeNameOf(v interface{}) string {
	return newTypeID(reflect.TypeOf(v))
}

func newTypeID(t reflect.Type) string {
	delimiter := "."
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		delimiter = ".*"
	}
	return t.PkgPath() + delimiter + t.Name()
}
