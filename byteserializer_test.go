package serializer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteSerializer(t *testing.T) {
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
	bs := NewByteSerializer(s)

	value := Object{
		12345,
		"foo",
	}
	data, err := bs.SerializeByte(value)
	assert.NoError(err)
	deserialized, err := bs.DeserializeByte(data)
	assert.NoError(err)
	assert.Equal(value, deserialized)

	pointer := &Object{
		54321,
		"bar",
	}
	data, err = bs.SerializeByte(pointer)
	assert.NoError(err)
	deserialized, err = bs.DeserializeByte(data)
	assert.NoError(err)
	assert.Equal(pointer, deserialized)
}
