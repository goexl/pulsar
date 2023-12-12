package builder

import (
	"github.com/goexl/pulsar/internal/core"
	"github.com/goexl/pulsar/internal/param"
)

type Client struct {
	params *param.Client
}

func NewClient() *Client {
	return &Client{
		params: param.NewClient(),
	}
}

func (c *Client) Server() *Server {
	return NewServer(c.params, c)
}

func (c *Client) Build() *core.Client {
	return core.NewClient(c.params)
}
