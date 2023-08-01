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

func (c *Client) Server(label string, url string) (client *Client) {
	c.param.Servers[label] = url
	client = c

	return
}

func (c *Client) Provider(label string, provider internal.Provider) (client *Client) {
	c.param.Providers[label] = provider
	client = c

	return
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
	label := connection.Label
	if url, uok := c.param.Servers[label]; !uok {
		err = exc.NewField("未能找到连接地址", field.New("label", connection.Label))
	} else if provider, ok := c.param.Providers[label]; !ok {
		err = exc.NewField("未能找到授权", field.New("label", connection.Label))
	} else {
		options.URL = url
		options.Authentication = pulsar.NewAuthenticationTokenFromSupplier(provider.Provide)
		client, err = pulsar.NewClient(options)
	}

	return
}
