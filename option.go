package serializer

type Option func(s *Serializer)

func WithMarshaler(m Marshaler) Option {
	return func(s *Serializer) {
		s.marshaler = m
	}
}

func WithEncoder(e Encoder) Option {
	return func(s *Serializer) {
		s.encoder = e
	}
}

func WithTypeNameResolver(r TypeNameResolver) Option {
	return func(s *Serializer) {
		s.resolver = r
	}
}
