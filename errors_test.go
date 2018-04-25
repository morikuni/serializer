package serializer

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	type Test struct {
		Input  error
		Expect string
	}

	tests := map[string]Test{
		"unknown type": {
			Input:  UnknownTypeError{"int"},
			Expect: `unknown type: "int": register it before serialize/deserialize`,
		},
		"duplicated name": {
			Input:  DuplicatedNameError{"int", reflect.TypeOf(int(123)), reflect.TypeOf(hello{})},
			Expect: `name duplicated: "int" and "serializer.hello" have same name: "int"`,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			assert.EqualError(t, test.Input, test.Expect)
		})
	}
}
