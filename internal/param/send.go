package param

import (
	"github.com/goexl/pulsar/internal/serializer"
)

type Send[T any] struct {
	*Producer[T]

	Key        string
	Properties map[string]string

	Encoder serializer.Encoder[T]
}

func NewSend[T any]() *Send[T] {
	return &Send[T]{
		Producer: NewProducer[T](),

		Encoder: serializer.NewJson[T](),
	}
}
