package builder

import (
	"github.com/goexl/pulsar/internal/core"
	"github.com/goexl/pulsar/internal/param"
)

type Producer[T any] struct {
	*Connection[T]

	param *param.Producer[T]
}

func NewProducer[T any]() *Producer[T] {
	return &Producer[T]{
		Connection: NewConnection[T](),

		param: param.NewProducer[T](),
	}
}

func (p *Producer[T]) Build() (*core.Producer[T], error) {
	return core.NewProducer[T]()
}
