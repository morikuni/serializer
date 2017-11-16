package serializer

import (
	"fmt"
)

type UnknownTypeError struct {
	TypeID TypeID
}

func (e UnknownTypeError) Error() string {
	return fmt.Sprintf("unknown type: %q", e.TypeID)
}
