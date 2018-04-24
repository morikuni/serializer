package serializer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoder(t *testing.T) {
	type Test struct {
		Input Encoder
	}

	tests := map[string]Test{
		"json": {
			Input: NewJSONEncoder(),
		},
		"text json": {
			Input: NewTextJSONEncoder(),
		},
		"protobuf": {
			Input: NewProtobufEncoder(),
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			e := test.Input

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
		})
	}
}
