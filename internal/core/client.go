package core

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/exc"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
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
	if cached, ok := c.clients[connection.Key()]; ok {
		client = cached.(pulsar.Client)
	} else {
		client, err = c.createClient(connection)
	}

	return
}

func (c *Client) createClient(connection *internal.Connection) (client pulsar.Client, err error) {
	options := pulsar.ClientOptions{}
	options.URL = c.param.Urls[connection.Label]
	options.Authentication = pulsar.NewAuthenticationTokenFromSupplier(c.param.Provider.Provide)
	if "" != options.URL {
		err = exc.NewField("未能找到连接地址", field.New("label", connection.Label))
	} else {
		client, err = pulsar.NewClient(options)
	}

	return
}
