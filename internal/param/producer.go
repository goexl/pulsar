package param

import (
	"github.com/goexl/pulsar/internal/internal"
)

type Producer[T any] struct {
	*internal.Connection
	*internal.Base[T]

	Topic string
}

func NewProducer[T any]() *Producer[T] {
	return &Producer[T]{
		Connection: internal.NewConnection(),
		Base:       internal.NewBase[T](),
	}
}
