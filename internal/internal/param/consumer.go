package param

import (
	"github.com/goexl/pulsar/internal/param"
)

type Consumer[T any] struct {
	*param.Connection
	*Base[T]

	Topics  []string
	Pattern string
}

func NewConsumer[T any]() *Consumer[T] {
	return &Consumer[T]{
		Connection: param.NewConnection(),
		Base:       NewBase[T](),
	}
}
