package serializer

type Option func(*serializer)

func WithTypeRegistry(tr TypeRegistry) Option {
	return func(s *serializer) {
		s.registry = tr
	}
}

func WithCodec(codec Codec) Option {
	return func(s *serializer) {
		s.codec = codec
	}
}
