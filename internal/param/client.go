package param

import (
	"time"
)

type Client struct {
	Queues map[string]*string
}

func NewClient() *Client {
	return &Client{
		Wait:   20 * time.Second,
		Queues: make(map[string]*string),
	}
}
