package param

import (
	"sync"

	"github.com/goexl/pulsar/internal/config"
)

type Client struct {
	Servers *sync.Map
}

func NewClient() *Client {
	return &Client{
		Servers: new(sync.Map),
	}
}

func (c *Client) Get(label string) (server *config.Server) {
	if cached, ok := c.Servers.Load(label); ok {
		server = cached.(*config.Server)
	} else {
		server = new(config.Server)
		c.Servers.Store(label, server)
	}

	return
}
