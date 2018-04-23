package serializer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeNameOf(t *testing.T) {
	type Test struct {
		Input  interface{}
		Expect string
	}

	tests := map[string]Test{
		"struct": {
			Input:  Data{},
			Expect: "github.com/morikuni/serializer.Data",
		},
		"struct-pointer": {
			Input:  &Data{},
			Expect: "github.com/morikuni/serializer.*Data",
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			assert.Equal(t, test.Expect, TypeNameOf(test.Input))
		})
	}
}
