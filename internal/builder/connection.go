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

func (b *Connection) Label(label string) (connection *Connection) {
	b.param.Label = label
	connection = b

	return
}

func (b *Connection) Name(name string) (connection *Connection) {
	b.param.Name = name
	connection = b

	return
}
