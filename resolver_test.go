package serializer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type hello struct{}

func (*hello) AliasName() []string {
	return []string{"world"}
}

func TestTypeNameResolver(t *testing.T) {
	type Expect struct {
		Name    string
		Aliases []string
	}
	type Test struct {
		Input  interface{}
		Expect Expect
	}

	i := 123
	tests := map[string]Test{
		"struct": {
			Input: Data{},
			Expect: Expect{
				"github.com/morikuni/serializer.Data",
				nil,
			},
		},
		"struct-pointer": {
			Input: &Data{},
			Expect: Expect{
				"github.com/morikuni/*serializer.Data",
				nil,
			},
		},
		"int": {
			Input: i,
			Expect: Expect{
				"int",
				nil,
			},
		},
		"int-pointer": {
			Input: &i,
			Expect: Expect{
				"*int",
				nil,
			},
		},
		"slice": {
			Input: []int{},
			Expect: Expect{
				"[]int",
				nil,
			},
		},
		"map": {
			Input: map[string]int{},
			Expect: Expect{
				"map[string]int",
				nil,
			},
		},
		"alias-namer": {
			Input: &hello{},
			Expect: Expect{
				"github.com/morikuni/*serializer.hello",
				[]string{"world"},
			},
		},
		"not-alias-namer": {
			Input: hello{},
			Expect: Expect{
				"github.com/morikuni/serializer.hello",
				nil,
			},
		},
	}

	r := NewTypeNameResolver()
	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			name, aliases := r.ResolveName(test.Input)
			assert.Equal(t, test.Expect.Name, name)
			assert.Equal(t, test.Expect.Aliases, aliases)
		})
	}
}
