package core

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/gox"
	"github.com/goexl/pulsar/internal/builder"
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/internal"
	"github.com/goexl/pulsar/internal/param"
)

type Producer[T any] struct {
	get   callback.GetClient
	param *param.Producer[T]

	_ gox.CannotCopy
}

func NewProducer[T any](get callback.GetClient, param *param.Producer[T]) *Producer[T] {
	return &Producer[T]{
		get:   get,
		param: param,
	}
}

func (p *Producer[T]) Send() *builder.Send[T] {
	return builder.NewSend[T](p.getProducer)
}

func (p *Producer[T]) getProducer(connection *internal.Connection[T]) (producer pulsar.Producer, err error) {
	return
}
