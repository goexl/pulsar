package builder

import (
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/internal/constant"
	"github.com/goexl/pulsar/internal/internal/param"
	"github.com/goexl/pulsar/internal/internal/worker"
)

type Sender[T any] struct {
	params     *param.Sender[T]
	client     callback.GetClient
	properties callback.GetProperties
}

func NewSender[T any](
	client callback.GetClient, server callback.GetServer, properties callback.GetProperties,
) *Sender[T] {
	return &Sender[T]{
		params:     param.NewSender[T](server),
		client:     client,
		properties: properties,
	}
}

func (s *Sender[T]) Tag(key string) (sender *Sender[T]) {
	s.params.Properties[key] = constant.Tags
	sender = s

	return
}

func (s *Sender[T]) Property(key string, value string) (sender *Sender[T]) {
	s.params.Properties[key] = value
	sender = s

	return
}

func (s *Sender[T]) Build() *worker.Sender[T] {
	return worker.NewSender[T](s.params, s.client, s.properties)
}
