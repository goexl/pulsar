package builder

import (
	"github.com/goexl/pulsar/internal/internal"
)

type Connection[T any] struct {
	param *internal.Connection[T]
}

func NewConnection[T any]() *Connection[T] {
	return &Connection[T]{
		param: internal.NewConnection[T](),
	}
}
