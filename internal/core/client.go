package core

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/exception"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/pulsar/internal/config"
	"github.com/goexl/pulsar/internal/internal/builder"
	"github.com/goexl/pulsar/internal/internal/core"
	"github.com/goexl/pulsar/internal/param"
)

type Client struct {
	clients map[string]pulsar.Client
	params  *param.Client
	_       gox.CannotCopy
}

func NewClient(param *param.Client) *Client {
	return &Client{
		clients: make(map[string]pulsar.Client),
		params:  param,
	}
}

func (c *Client) Server(label string, url string) (client *Client) {
	c.params.Servers.Store(label, url)
	client = c

	return
}

func (c *Client) Provider(label string, provider core.Provider) (client *Client) {
	if cached, ok := c.params.Servers.Load(label); ok {
		cached.(*config.Server).Provider = provider
	}
	client = c

	return
}

func (c *Client) Consumer() *builder.Consumer[any] {
	return builder.NewConsumer[any](c.getClient)
}

func (c *Client) Producer(topic string) *builder.Producer[any] {
	return builder.NewProducer[any](topic)
}

func (c *Client) getClient(connection *param.Connection) (client pulsar.Client, err error) {
	if cached, ok := c.clients[connection.Key()]; ok {
		client = cached.(pulsar.Client)
	} else {
		client, err = c.createClient(connection)
	}

	return
}

func (c *Client) createClient(connection *param.Connection) (client pulsar.Client, err error) {
	options := pulsar.ClientOptions{}
	label := connection.Label
	if url, uok := c.params.Servers.Load(label); !uok {
		err = exception.New().Message("未能找到连接地址").Field(field.New("label", connection.Label)).Build()
	} else if provider, pok := c.params.Providers[label]; !pok {
		err = exception.New().Message("未能找到授权").Field(field.New("label", connection.Label)).Build()
	} else {
		options.URL = url
		options.Authentication = pulsar.NewAuthenticationTokenFromSupplier(provider.Provide)
		client, err = pulsar.NewClient(options)
	}

	return
}
