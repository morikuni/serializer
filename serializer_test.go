package serializer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializer(t *testing.T) {
	type Object struct {
		ID   int64
		Name string
		Age  *int
	}

	x := 321
	tests := map[string]interface{}{
		"struct":         Object{12345, "aaaa", &x},
		"struct-pointer": &Object{54321, "bbbb", nil},
		"int":            123,
		"int-pointer":    &x,
	}

	s := NewSerializer(
		NewJSONMarshaler(),
		NewJSONEncoder(),
	)
	s.Register(
		Object{},
		(*Object)(nil),
		int(123),
		(*int)(nil),
	)

	for name, input := range tests {
		t.Run(name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			err := s.Serialize(buf, input)
			assert.NoError(t, err)
			v, err := s.Deserialize(buf)
			assert.NoError(t, err)
			assert.Equal(t, input, v)
		})
	}
}
