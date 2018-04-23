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
	return fmt.Sprintf("unknown type: %q", e.Name)
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
