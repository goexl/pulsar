package builder

import (
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/core"
	param2 "github.com/goexl/pulsar/internal/internal/param"
)

type Consumer[T any] struct {
	*Connection
	*param2.Base[T]

	param *param2.Consumer[T]
	get   callback.GetClient
}

func NewConsumer[T any](get callback.GetClient) *Consumer[T] {
	return &Consumer[T]{
		Connection: NewConnection(),
		Base:       param2.NewBase[T](),

		param: param2.NewConsumer[T](),
		get:   get,
	}
}

func (c *Consumer[T]) Build() *core.Consumer[T] {
	return core.NewConsumer[T](c.param, c.get)
}
