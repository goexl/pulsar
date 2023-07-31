package param

import (
	"github.com/goexl/pulsar/internal/config"
	"github.com/goexl/pulsar/internal/internal"
)

type Client struct {
	Urls     map[string]string
	Region   string
	Timeout  config.Timeout
	Provider internal.Provider
}

func NewClient() *Client {
	return &Client{
		Urls: make(map[string]string),
	}
}
