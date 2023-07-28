package builder

import (
	"github.com/goexl/pulsar/internal/internal"
)

type Connection struct {
	param *internal.Connection
}

func NewConnection() *Connection {
	return &Connection{
		param: internal.NewConnection(),
	}
}
