package serializer

import (
	"fmt"
)

// UnknownTypeError is a error used when the type
// is not a registered into a Serializer.
type UnknownTypeError struct {
	Name string
}

// Error implements error.
func (e UnknownTypeError) Error() string {
	return fmt.Sprintf("unknown type: %q: you should call Serializer.Register", e.Name)
}

// UnsupportedTypeError is a error used when the type
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

type DuplicatedNameError struct {
	Name string
	Old  string
	New  string
}

// Error implements error.
func (e DuplicatedNameError) Error() string {
	return fmt.Sprintf("name duplicated: %s: %s has registered but %s has same name", e.Name, e.Old, e.New)
}
