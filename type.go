package serializer

import (
	"reflect"
)

// TypeNameOf returns a unique name of the
// type of v.
func TypeNameOf(v interface{}) string {
	t := reflect.TypeOf(v)
	delimiter := "."
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		delimiter = ".*"
	}
	return t.PkgPath() + delimiter + t.Name()
}
