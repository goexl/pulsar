package param

import (
	"github.com/goexl/pulsar/internal/internal"
)

type Producer[T any] struct {
	*internal.Connection[T]
}

func NewProducer[T any]() *Producer[T] {
	return &Producer[T]{
		Connection: internal.NewConnection[T](),
	}
}
