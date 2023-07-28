package builder

import (
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/param"
	"github.com/goexl/pulsar/internal/serializer"
	"github.com/goexl/pulsar/internal/worker"
)

type Send[T any] struct {
	*Base

	param *param.Send[T]
	send  callback.Send
}

func NewSend[T any](send callback.Send) *Send[T] {
	return &Send[T]{
		Base: NewBase(),

		param: param.NewSend[T](send),
		send:  send,
	}
}

func (s *Send[T]) Encoder(encoder serializer.Encoder[T]) (send *Send[T]) {
	s.param.Encoder = encoder
	send = s

	return
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

func (s *Send[T]) Build() *worker.Send {
	return worker.NewSend(s.param)
}
