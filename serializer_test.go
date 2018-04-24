package serializer

import (
	"bytes"
	"io/ioutil"
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

func BenchmarkSerializer_Serialize(b *testing.B) {
	s := New()
	assert.NoError(b, s.Register(
		(*Data)(nil),
	))

	h := &Data{"aaa", []byte{1, 2, 3}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Serialize(ioutil.Discard, h)
	}
}

func BenchmarkSerializer_Deserialize(b *testing.B) {
	s := New()
	assert.NoError(b, s.Register(
		(*Data)(nil),
	))

	b.ResetTimer()

	buf := &bytes.Buffer{}
	assert.NoError(b, s.Serialize(buf, &Data{"aaa", []byte{1, 2, 3}}))

	r := bytes.NewReader(buf.Bytes())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Deserialize(r)
	}
}
