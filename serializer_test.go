package serializer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializer(t *testing.T) {
	type Object struct {
		ID        int64
		Name      string
		Age       *int
		Nicknames []string
	}

	x := 321
	tests := map[string]interface{}{
		"struct":         Object{12345, "aaaa", &x, []string{"a", "b"}},
		"struct pointer": &Object{54321, "bbbb", nil, nil},
		"int":            123,
		"int pointer":    &x,
		"slice":          []int32{1, 2, 3},
		"map":            map[string]int{"one": 1, "two": 2},
		"alias":          &hello{3},
	}

	s := New(
		WithMarshaler(NewJSONMarshaler()),
		WithEncoder(NewJSONEncoder()),
	)
	assert.NoError(t, s.Register(
		Object{},
		(*Object)(nil),
		int(123),
		(*int)(nil),
		[]int32{},
		map[string]int{},
		(*hello)(nil),
	))

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
