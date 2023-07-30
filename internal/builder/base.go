package builder

import (
	"github.com/goexl/pulsar/internal/internal"
	"github.com/goexl/pulsar/internal/serializer"
)

type Base[T any] struct {
	param *internal.Base[T]
}

func NewBase[T any]() *Base[T] {
	return &Base[T]{
		param: internal.NewBase[T](),
	}
}

func (s *Send[T]) Encoder(encoder serializer.Encoder[T]) (send *Send[T]) {
	s.param.Encoder = encoder
	send = s

	return
}

func (b *Base[T]) Label(label string) (base *Base[T]) {
	b.param.Label = label
	base = b

	return
}
