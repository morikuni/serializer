package serializer

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewType(t *testing.T) {

	type Object struct {
		Name string
	}

	type Input struct {
		Value interface{}
	}
	type Expect struct {
		Type Type
	}
	type Test struct {
		Title  string
		Input  Input
		Expect Expect
	}

	tests := []Test{
		{
			Title: "struct",
			Input: Input{Object{}},
			Expect: Expect{
				Type{
					"github.com/morikuni/serializer.Object",
					constructor{
						false,
						reflect.TypeOf(Object{}),
					},
				},
			},
		},
		{
			Title: "pointer",
			Input: Input{(*Object)(nil)},
			Expect: Expect{
				Type{
					"github.com/morikuni/serializer.*Object",
					constructor{
						true,
						reflect.TypeOf(Object{}),
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(test.Expect.Type, NewType(test.Input.Value))
		})
	}
}

func TestConstructor(t *testing.T) {
	assert := assert.New(t)

	type Object struct {
		Name string
	}

	cv := NewConstructor(Object{}).(constructor)
	assert.Equal(&Object{}, cv.newPointer())
	assert.Equal(Object{}, cv.restore(cv.newPointer()))

	cp := NewConstructor((*Object)(nil)).(constructor)
	assert.Equal(&Object{}, cp.newPointer())
	assert.Equal(&Object{}, cp.restore(cp.newPointer()))
}
