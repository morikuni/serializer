package serializer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeNameResolver(t *testing.T) {
	type Test struct {
		Input  interface{}
		Expect string
	}

	i := 123
	tests := map[string]Test{
		"struct": {
			Input:  Data{},
			Expect: "github.com/morikuni/serializer.Data",
		},
		"struct-pointer": {
			Input:  &Data{},
			Expect: "github.com/morikuni/*serializer.Data",
		},
		"int": {
			Input:  i,
			Expect: "int",
		},
		"int-pointer": {
			Input:  &i,
			Expect: "*int",
		},
		"slice": {
			Input:  []int{},
			Expect: "[]int",
		},
		"map": {
			Input:  map[string]int{},
			Expect: "map[string]int",
		},
	}

	r := NewTypeNameResolver()
	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			assert.Equal(t, test.Expect, r.ResolveName(test.Input))
		})
	}
}
