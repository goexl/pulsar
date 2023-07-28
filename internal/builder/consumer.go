package builder

import (
	"github.com/goexl/pulsar/internal/core"
)

type Consumer struct {
	*Connection
}

func NewConsumer() *Consumer {
	return &Consumer{
		Connection: NewConnection(),
	}
}

func (c *Consumer) Build() *core.Consumer {
	return core.NewConsumer()
}
