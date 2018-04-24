package serializer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshaler(t *testing.T) {
	type Test struct {
		Input Marshaler
	}

	tests := map[string]Test{
		"json": {
			Input: NewJSONMarshaler(),
		},
		"protobuf": {
			Input: NewProtobufMarshaler(),
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			e := test.Input

			d1 := Data{
				Name:    "aaa",
				Payload: []byte{1, 2, 3},
			}

			data, err := e.Marshal(&d1)
			assert.NoError(t, err)

			var d2 Data
			err = e.Unmarshal(data, &d2)
			assert.NoError(t, err)
			assert.Equal(t, d1, d2)
		})
	}
}
