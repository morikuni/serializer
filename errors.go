package serializer

import (
	"fmt"
	"reflect"
)

// UnknownTypeError is an error used when the type
// is not a registered into a Serializer.
type UnknownTypeError struct {
	Name string
}

// Error implements error.
func (e UnknownTypeError) Error() string {
	return fmt.Sprintf("unknown type: %q: register it before serialize/deserialize", e.Name)
}

// UnsupportedTypeError is an error used when the type
// is not supported by a marshaler. (e.g. the type did not
// implement a protocol buffer interface.)
type UnsupportedTypeError struct {
	Name   string
	Reason string
}

// Error implements error.
func (e UnsupportedTypeError) Error() string {
	return fmt.Sprintf("unsupported type: %q: %s", e.Name, e.Reason)
}

// DuplicatedNameError is an error used when the name
// is associated with more than one type.
type DuplicatedNameError struct {
	Name string
	Old  reflect.Type
	New  reflect.Type
}

// Error implements error.
func (e DuplicatedNameError) Error() string {
	return fmt.Sprintf("name duplicated: %q and %q have same name: %q", e.Old.String(), e.New.String(), e.Name)
}
