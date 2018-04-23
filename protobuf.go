package serializer

import (
	"io"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

func NewProtobufMarshaler() Marshaler {
	return protobufMarshaler{}
}

type protobufMarshaler struct{}

var _ interface {
	Marshaler
} = protobufMarshaler{}

func (protobufMarshaler) Marshal(v interface{}) ([]byte, error) {
	msg, ok := v.(proto.Message)
	if !ok {
		return nil, UnsupportedTypeError{TypeNameOf(v), "does not implement proto.Message"}
	}
	return proto.Marshal(msg)
}

func (protobufMarshaler) Unmarshal(data []byte, v interface{}) error {
	msg, ok := v.(proto.Message)
	if !ok {
		return UnsupportedTypeError{TypeNameOf(v), "does not implement proto.Message"}
	}
	return proto.Unmarshal(data, msg)
}

func NewProtobufEncoder() Encoder {
	return protobufEncoder{}
}

type protobufEncoder struct{}

func (protobufEncoder) Encode(w io.Writer, d Data) error {
	msg, err := proto.Marshal(&d)
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	return err
}

func (protobufEncoder) Decode(r io.Reader) (Data, error) {
	msg, err := ioutil.ReadAll(r)
	if err != nil {
		return Data{}, err
	}
	var d Data
	if err := proto.Unmarshal(msg, &d); err != nil {
		return Data{}, err
	}
	return d, nil
}
