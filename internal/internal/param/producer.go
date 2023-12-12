package param

import (
	"github.com/goexl/pulsar/internal/param"
)

type Producer[T any] struct {
	*param.Connection
	*Base[T]

	Topic string
}

func NewProducer[T any]() *Producer[T] {
	return &Producer[T]{
		Connection: param.NewConnection(),
		Base:       NewBase[T](),
	}
}
