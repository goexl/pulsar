package core

import (
	"sync"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/gox"
	"github.com/goexl/pulsar/internal/builder"
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/param"
)

type Consumer[T any] struct {
	get       callback.GetClient
	param     *param.Consumer[T]
	consumers *sync.Map

	_ gox.CannotCopy
}

func NewConsumer[T any](param *param.Consumer[T], get callback.GetClient) *Consumer[T] {
	return &Consumer[T]{
		param:     param,
		get:       get,
		consumers: new(sync.Map),
	}
}

func (c *Consumer[T]) Handle() *builder.Handle[T] {
	return builder.NewHandle[T](c.getConsumer)
}

func (c *Consumer[T]) getConsumer(config *param.Consumer[T]) (consumer pulsar.Consumer, err error) {
	if client, ge := c.get(config.Connection); nil != ge {
		err = ge
	} else if cached, ok := c.consumers.Load(config.Key()); ok {
		consumer = cached.(pulsar.Consumer)
	} else {
		consumer, err = c.createConsumer(client, config)
	}

	return
}

func (c *Consumer[T]) createConsumer(
	client pulsar.Client,
	config *param.Consumer[T],
) (consumer pulsar.Consumer, err error) {
	options := pulsar.ConsumerOptions{}
	options.Name = config.Name
	options.Topics = config.Topics
	options.TopicsPattern = config.Pattern
	consumer, err = client.Subscribe(options)

	return
}
