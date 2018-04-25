package serializer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

type a struct {
	Int          int
	Int32Pointer *int32
	String       string
	Slice        []int64
	Map          map[string]float64
}

func TestJSONEncoder(t *testing.T) {
	type Test struct {
		Input a
	}

	i32 := int32(321)
	tests := map[string]Test{
		"normal": Test{
			a{
				123,
				&i32,
				"456あああ",
				[]int64{1, 2, 3},
				map[string]float64{
					"negative": -7.89,
				},
			},
		},
		"zero value": Test{
			a{
				0,
				nil,
				"",
				nil,
				nil,
			},
		},
	}

	e := NewJSONEncoder()
	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			buf := &bytes.Buffer{}
			err := e.Encode(buf, "aaa", test.Input)
			assert.NoError(t, err)

			id, payload, err := e.Decode(buf)
			assert.NoError(t, err)
			assert.Equal(t, "aaa", id)

			var r a
			err = e.Unmarshal(payload, &r)
			assert.NoError(t, err)
			assert.Equal(t, test.Input, r)
		})
	}
}
