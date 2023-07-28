package param

import (
	"github.com/goexl/pulsar/internal/internal"
)

type Producer struct {
	*internal.Connection
}

func NewProducer() *Producer {
	return &Producer{
		Connection: internal.NewConnection(),
	}
}
