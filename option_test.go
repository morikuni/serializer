package serializer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOption(t *testing.T) {
	s := Serializer{}

	assert.Nil(t, s.marshaler)
	assert.Nil(t, s.encoder)

	m := NewProtobufMarshaler()
	e := NewTextJSONEncoder()
	WithMarshaler(m)(&s)
	WithEncoder(e)(&s)

	assert.Equal(t, m, s.marshaler)
	assert.Equal(t, e, s.encoder)
}
