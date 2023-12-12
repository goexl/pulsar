package param

import (
	"github.com/goexl/exception"
	"github.com/goexl/log"
	"github.com/goexl/pulsar/internal/internal/constant"
)

type Client struct {
	Servers map[string]*Server
	Logger  log.Logger
}

func NewClient() *Client {
	return &Client{
		Servers: make(map[string]*Server),
		Logger:  log.New().Apply(),
	}
}

func (c *Client) Get(label string) (server *Server, err error) {
	if cached, ok := c.Servers[label]; ok {
		server = cached
	} else {
		err = exception.New().Message("未找到服务器").Build()
	}

	return
}

func (c *Client) ProducerProperties(label string) (properties map[string]string) {
	properties = make(map[string]string)
	if server, ok := c.Servers[label]; ok {
		for key, value := range server.Properties {
			properties[key] = value
		}
		for _, tag := range server.Tags {
			properties[tag] = constant.Tags
		}
		for key, value := range server.Producer.Properties {
			properties[key] = value
		}
		for _, tag := range server.Producer.Tags {
			properties[tag] = constant.Tags
		}
	}

	return
}

func (c *Client) ConsumerProperties(label string) (properties map[string]string) {
	properties = make(map[string]string)
	if server, ok := c.Servers[label]; ok {
		for key, value := range server.Properties {
			properties[key] = value
		}
		for _, tag := range server.Tags {
			properties[tag] = constant.Tags
		}
		for key, value := range server.Consumer.Properties {
			properties[key] = value
		}
		for _, tag := range server.Consumer.Tags {
			properties[tag] = constant.Tags
		}
	}

	return
}
