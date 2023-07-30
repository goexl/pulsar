package param

import (
	"github.com/goexl/pulsar/internal/internal"
	"github.com/goexl/pulsar/internal/serializer"
)

type Send[T any] struct {
	*internal.Base[T]

	Key        string
	Properties map[string]string

	Encoder serializer.Encoder[T]
}

func NewSend[T any]() *Send[T] {
	return &Send[T]{
		Base: internal.NewBase[T](),

		Encoder: serializer.NewJson[T](),
	}
}
