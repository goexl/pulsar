package core

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/gox"
	"github.com/goexl/pulsar/internal/internal"
	"github.com/goexl/pulsar/internal/param"
)

type Client struct {
	clients map[string]pulsar.Client
	param   *param.Client

	_ gox.CannotCopy
}

func NewClient(param *param.Client) *Client {
	return &Client{
		clients: make(map[string]pulsar.Client),
		param:   param,
	}
}

func (c *Client) Consumer() *ConsumerBuilder[any] {
	return NewConsumerBuilder[any](c.getClient)
}

func (c *Client) Producer(topic string) *ProducerBuilder[any] {
	return NewProducerBuilder[any](topic)
}

func (c *Client) getClient(connection *internal.Connection) (client pulsar.Client, err error) {
	return
}
