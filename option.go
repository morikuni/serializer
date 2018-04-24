package serializer

// Option represents a option of the serializer.
type Option func(s *Serializer)

// WithMarshaler changes the marshaler.
func WithMarshaler(m Marshaler) Option {
	return func(s *Serializer) {
		s.marshaler = m
	}
}

// WithEncoder changes the encoder.
func WithEncoder(e Encoder) Option {
	return func(s *Serializer) {
		s.encoder = e
	}
}

// WithTypeNameResolver changes the resolver.
func WithTypeNameResolver(r TypeNameResolver) Option {
	return func(s *Serializer) {
		s.resolver = r
	}
}
