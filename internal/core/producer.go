package core

import (
	"sync"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/gox"
	"github.com/goexl/pulsar/internal/builder"
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/param"
)

type Producer[T any] struct {
	get       callback.GetClient
	param     *param.Producer[T]
	producers *sync.Map

	_ gox.CannotCopy
}

func NewProducer[T any](get callback.GetClient, param *param.Producer[T]) *Producer[T] {
	return &Producer[T]{
		get:       get,
		param:     param,
		producers: new(sync.Map),
	}
}

func (p *Producer[T]) Send() *builder.Send[T] {
	return builder.NewSend[T](p.getProducer)
}

func (p *Producer[T]) getProducer(config *param.Producer[T]) (producer pulsar.Producer, err error) {
	if client, ge := p.get(config.Connection); nil != ge {
		err = ge
	} else if cached, ok := p.producers.Load(config.Key()); ok {
		producer = cached.(pulsar.Producer)
	} else {
		producer, err = p.createProducer(client, config)
	}

	return
}

func (p *Producer[T]) createProducer(client pulsar.Client, config *param.Producer[T]) (producer pulsar.Producer, err error) {
	options := pulsar.ProducerOptions{}
	options.Name = config.Name
	options.Topic = config.Topic
	producer, err = client.CreateProducer(options)

	return
}
