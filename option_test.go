package serializer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOption(t *testing.T) {
	s := Serializer{}

	assert.Nil(t, s.encoder)

	e := NewJSONEncoder()
	r := NewTypeNameResolver()
	WithEncoder(e)(&s)
	WithTypeNameResolver(r)(&s)

	assert.Equal(t, e, s.encoder)
	assert.Equal(t, r, s.resolver)
}
