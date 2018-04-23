package serializer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONMarshaler(t *testing.T) {
	m := NewJSONMarshaler()

	d1 := Data{
		Name:    "aaa",
		Payload: []byte{1, 2, 3},
	}

	data, err := m.Marshal(&d1)
	assert.NoError(t, err)

	var d2 Data
	assert.NoError(t, m.Unmarshal(data, &d2))
	assert.Equal(t, d1, d2)
}

func TestJSONEncoder(t *testing.T) {
	e := NewJSONEncoder()

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

func TestTextJSONEncoder(t *testing.T) {
	e := NewTextJSONEncoder()

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
