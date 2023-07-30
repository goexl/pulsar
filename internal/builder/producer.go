package builder

import (
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/core"
	"github.com/goexl/pulsar/internal/param"
)

type Producer[T any] struct {
	*Connection[T]

	param *param.Producer[T]
	get   callback.GetClient
}

func NewProducer[T any](topic string) (producer *Producer[T]) {
	producer = new(Producer[T])
	producer.Connection = NewConnection[T]()
	producer.param = param.NewProducer[T]()

	producer.Connection.param.Topic = topic

	return
}

func (p *Producer[T]) Build() *core.Producer[T] {
	return core.NewProducer[T](p.get, p.param)
}
