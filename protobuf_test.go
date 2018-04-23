package serializer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProtobufMarshaler(t *testing.T) {
	m := NewProtobufMarshaler()

	d1 := Data{
		Name:    "aaa",
		Payload: []byte{1, 2, 3},
	}

	data, err := m.Marshal(&d1)
	assert.NoError(t, err)

	var d2 Data
	assert.NoError(t, m.Unmarshal(data, &d2))
	assert.Equal(t, d1, d2)

	var i int
	_, err = m.Marshal(&i)
	assert.Error(t, err)
	assert.IsType(t, UnsupportedTypeError{}, err)
}

func TestProtobufEncoder(t *testing.T) {
	e := NewProtobufEncoder()

	d1 := Data{
		Name:    "aaa",
		Payload: []byte{1, 2, 3},
	}

	buf := &bytes.Buffer{}

	err := e.Encode(buf, d1)
	assert.NoError(t, err)

	d2, err := e.Decode(buf)
	assert.NoError(t, err)
	assert.Equal(t, d1, d2)
}
