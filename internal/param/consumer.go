package param

import (
	"github.com/goexl/pulsar/internal/internal"
)

type Consumer[T any] struct {
	*internal.Connection[T]
}

func NewConsumer[T any]() *Consumer[T] {
	return &Consumer[T]{
		Connection: internal.NewConnection[T](),
	}
}
