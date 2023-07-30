package core

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/gox"
	"github.com/goexl/pulsar/internal/builder"
	"github.com/goexl/pulsar/internal/param"
)

type Client struct {
	pulsars map[string]pulsar.Client
	param   *param.Client

	_ gox.CannotCopy
}

func NewClient(param *param.Client) *Client {
	return &Client{
		pulsars: make(map[string]pulsar.Client),
		param:   param,
	}
}

func (c *Client) Consumer(topic string) *builder.Consumer[any] {
	return builder.NewConsumer[any](topic)
}

func (c *Client) Producer(topic string) *builder.Producer[any] {
	return builder.NewProducer[any](topic)
}
