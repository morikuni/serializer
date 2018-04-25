package serializer

// Option represents a option of the serializer.
type Option func(s *Serializer)

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
