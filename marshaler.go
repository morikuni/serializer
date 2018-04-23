package serializer

// Marshaler marshals and unmarshals a object.
type Marshaler interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}
