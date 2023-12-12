package param

import (
	"github.com/goexl/pulsar/internal/serializer"
)

type Base[T any] struct {
	Encoder serializer.Encoder[T]
	Decoder serializer.Decoder[T]
}

func NewBase[T any]() *Base[T] {
	return &Base[T]{
		Encoder: serializer.NewJson[T](),
	}
}
