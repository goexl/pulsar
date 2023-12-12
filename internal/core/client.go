package core

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/exception"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/pulsar/internal/internal/builder"
	"github.com/goexl/pulsar/internal/param"
)

type Client struct {
	clients map[string]pulsar.Client
	params  *param.Client
	_       gox.CannotCopy
}

func NewClient(params *param.Client) *Client {
	return &Client{
		clients: make(map[string]pulsar.Client),
		params:  params,
	}
}

func (c *Client) Sender() *builder.Sender[any] {
	return builder.NewSender[any](c.get, c.params.Get, c.params.ProducerProperties)
}

func (c *Client) Handler() *builder.Handler[any] {
	return builder.NewHandler[any](c.get, c.params.Get, c.params.ConsumerProperties, c.params.Logger)
}

func (c *Client) get(label string) (client pulsar.Client, err error) {
	if cached, ok := c.clients[label]; ok {
		client = cached.(pulsar.Client)
	} else {
		client, err = c.create(label)
	}

	return
}

func (c *Client) create(label string) (client pulsar.Client, err error) {
	options := pulsar.ClientOptions{}
	if server, uok := c.params.Servers[label]; !uok {
		err = exception.New().Message("未能找到连接地址").Field(field.New("label", label)).Build()
	} else if nil == server.Provider {
		err = exception.New().Message("未能找到授权").Field(field.New("label", label)).Build()
	} else {
		options.URL = server.Url
		options.Authentication = pulsar.NewAuthenticationTokenFromSupplier(server.Provider.Provide)
		client, err = pulsar.NewClient(options)
	}

	return
}
