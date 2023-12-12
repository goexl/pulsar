package builder

import (
	"github.com/goexl/pulsar/internal/internal/param"
	"github.com/goexl/pulsar/internal/serializer"
)

type Base[T any] struct {
	param *param.Base[T]
}

func NewBase[T any]() *Base[T] {
	return &Base[T]{
		param: param.NewBase[T](),
	}
}

func (b *Base[T]) Encoder(encoder serializer.Encoder[T]) (base *Base[T]) {
	b.param.Encoder = encoder
	base = b

	return
}

func (b *Base[T]) Decoder(decoder serializer.Decoder[T]) (base *Base[T]) {
	b.param.Decoder = decoder
	base = b

	return
}

func (b *Base[T]) Serializer(serializer serializer.Serializer[T]) (base *Base[T]) {
	b.param.Encoder = serializer
	b.param.Decoder = serializer
	base = b

	return
}
