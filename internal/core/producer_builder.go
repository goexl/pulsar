package core

import (
	"github.com/goexl/pulsar/internal/builder"
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/param"
)

type ProducerBuilder[T any] struct {
	*builder.Connection
	*builder.Base[T]

	param *param.Producer[T]
	get   callback.GetClient
}

func NewProducerBuilder[T any](topic string) (producer *ProducerBuilder[T]) {
	producer = new(ProducerBuilder[T])
	producer.Connection = builder.NewConnection()
	producer.Base = builder.NewBase[T]()
	producer.param = param.NewProducer[T]()

	producer.param.Topic = topic

	return
}

func (p *ProducerBuilder[T]) Build() *Producer[T] {
	return NewProducer[T](p.get, p.param)
}
