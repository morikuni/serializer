package serializer

import "testing"

func TestData(t *testing.T) {
	// call protobuf methods for test coverage.
	var d Data
	d.Reset()
	d.String()
	d.ProtoMessage()
	d.Descriptor()
	d.GetName()
	d.GetPayload()
}
