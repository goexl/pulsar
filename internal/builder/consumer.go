package builder

import (
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/core"
	"github.com/goexl/pulsar/internal/param"
)

type Consumer[T any] struct {
	*Connection[T]

	param *param.Consumer[T]
	get   callback.GetClient
}

func NewConsumer[T any](topic string) (consumer *Consumer[T]) {
	consumer = new(Consumer[T])
	consumer.Connection = NewConnection[T]()
	consumer.param = param.NewConsumer[T]()

	consumer.Connection.param.Topic = topic

	return
}

func (c *Consumer[T]) Build() *core.Consumer[T] {
	return core.NewConsumer[T]()
}
