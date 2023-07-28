package internal

import (
	"github.com/goexl/pulsar/internal/serializer"
)

type Base[T any] struct {
	Serializer serializer.Serializer[T]
	Topic      string
	Name       string
}

func NewBase[T any]() *Base[T] {
	return &Base[T]{
		Serializer: serializer.NewJson[T](),
	}
}
