package builder

import (
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/param"
	"github.com/goexl/pulsar/internal/worker"
)

type Send[T any] struct {
	*Base[T]

	param *param.Send[T]
	get   callback.GetProducer[T]
}

func NewSend[T any](get callback.GetProducer[T]) *Send[T] {
	return &Send[T]{
		Base: NewBase[T](),

		param: param.NewSend[T](),
		get:   get,
	}
}

func (s *Send[T]) Key(key string) (send *Send[T]) {
	s.param.Key = key
	send = s

	return
}

func (s *Send[T]) Property(key string, value string) (send *Send[T]) {
	s.param.Properties[key] = value
	send = s

	return
}

func (s *Send[T]) Build() *worker.Send[T] {
	return worker.NewSend[T](s.param, s.get)
}
