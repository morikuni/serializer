package serializer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializer(t *testing.T) {
	assert := assert.New(t)

	type Object struct {
		ID   int64
		Name string
	}

	codec := NewJSONCodec()
	s := NewSerializer(WithCodec(codec))
	s.Register(
		Object{},
		(*Object)(nil),
	)

	value := Object{
		12345,
		"foo",
	}
	buf := &bytes.Buffer{}
	assert.NoError(s.Serialize(buf, value))
	deserialized, err := s.Deserialize(buf)
	assert.NoError(err)
	assert.Equal(value, deserialized)

	pointer := &Object{
		54321,
		"bar",
	}
	buf = &bytes.Buffer{}
	assert.NoError(s.Serialize(buf, pointer))
	deserialized, err = s.Deserialize(buf)
	assert.NoError(err)
	assert.Equal(pointer, deserialized)
}
