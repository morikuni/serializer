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
		"unsupported type": {
			Input:  UnsupportedTypeError{"int", "because test"},
			Expect: `unsupported type: "int": because test`,
		},
		"duplicated name": {
			Input:  DuplicatedNameError{"int", reflect.TypeOf(int(123)), reflect.TypeOf(Data{})},
			Expect: `name duplicated: "int" and "serializer.Data" have same name: "int"`,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			assert.EqualError(t, test.Input, test.Expect)
		})
	}
}
