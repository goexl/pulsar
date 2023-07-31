package core

import (
	"github.com/goexl/pulsar/internal/builder"
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/internal"
	"github.com/goexl/pulsar/internal/param"
)

type ConsumerBuilder[T any] struct {
	*builder.Connection
	*internal.Base[T]

	param *param.Consumer[T]
	get   callback.GetClient
}

func NewConsumerBuilder[T any](get callback.GetClient) *ConsumerBuilder[T] {
	return &ConsumerBuilder[T]{
		Connection: builder.NewConnection(),
		Base:       internal.NewBase[T](),

		param: param.NewConsumer[T](),
		get:   get,
	}
}

func (c *ConsumerBuilder[T]) Build() *Consumer[T] {
	return NewConsumer[T](c.param, c.get)
}
