package param

import (
	"github.com/goexl/pulsar/internal/internal"
)

type Consumer[T any] struct {
	*internal.Connection
	*internal.Base[T]

	Topics  []string
	Pattern string
}

func NewConsumer[T any]() *Consumer[T] {
	return &Consumer[T]{
		Connection: internal.NewConnection(),
		Base:       internal.NewBase[T](),
	}
}
