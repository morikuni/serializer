package serializer

import "bytes"

type ByteSerializer interface {
	SerializeByte(v interface{}) ([]byte, error)
	DeserializeByte(data []byte) (interface{}, error)
}

func NewByteSerializer(serializer Serializer) ByteSerializer {
	return byteSerializer{serializer}
}

type byteSerializer struct {
	serializer Serializer
}

func (bs byteSerializer) SerializeByte(v interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := bs.serializer.Serialize(buf, v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (bs byteSerializer) DeserializeByte(data []byte) (interface{}, error) {
	buf := bytes.NewReader(data)
	return bs.serializer.Deserialize(buf)
}
