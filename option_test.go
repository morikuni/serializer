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
	r := NewTypeNameResolver()
	WithMarshaler(m)(&s)
	WithEncoder(e)(&s)
	WithTypeNameResolver(r)(&s)

	assert.Equal(t, m, s.marshaler)
	assert.Equal(t, e, s.encoder)
	assert.Equal(t, r, s.resolver)
}
