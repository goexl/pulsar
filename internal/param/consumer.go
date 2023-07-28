package param

import (
	"github.com/goexl/pulsar/internal/internal"
)

type Consumer struct {
	*internal.Connection
}

func NewConsumer() *Consumer {
	return &Consumer{
		Connection: internal.NewConnection(),
	}
}
