package serializer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProtobuf(t *testing.T) {
	m := NewProtobufMarshaler()

	_, err := m.Marshal(123)
	assert.IsType(t, UnsupportedTypeError{}, err)
	err = m.Unmarshal(nil, 123)
	assert.IsType(t, UnsupportedTypeError{}, err)
}
