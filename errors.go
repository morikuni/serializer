package serializer

import (
	"fmt"
)

type UnknownTypeError struct {
	Name string
}

func (e UnknownTypeError) Error() string {
	return fmt.Sprintf("unknown type: %q", e.Name)
}

type UnsupportedTypeError struct {
	Name   string
	Reason string
}

func (e UnsupportedTypeError) Error() string {
	return fmt.Sprintf("unsupported type: %q: %s", e.Name, e.Reason)
}
