package core

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/gox"
	"github.com/goexl/pulsar/internal/builder"
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/internal"
	"github.com/goexl/pulsar/internal/param"
)

type Consumer[T any] struct {
	get   callback.GetClient
	param *param.Consumer[T]

	_ gox.CannotCopy
}

func NewConsumer[T any](get callback.GetClient, param *param.Consumer[T]) *Consumer[T] {
	return &Consumer[T]{
		get:   get,
		param: param,
	}
}

func (c *Consumer[T]) Handle(topic string) *builder.Handle[T] {
	return builder.NewHandle[T](c.getConsumer)
}

func (c *Consumer[T]) getConsumer(connection *internal.Connection[T]) (consumer pulsar.Consumer, err error) {
	return
}
