package serializer

import (
	"reflect"
)

// TypeNameOf returns a unique name of the
// type of v.
func TypeNameOf(v interface{}) string {
	t := reflect.TypeOf(v)

	ptr := ""
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		ptr = "*"
	}

	pkg := t.PkgPath() + "."
	if pkg == "." {
		pkg = ""
	}

	return pkg + ptr + t.Name()
}
