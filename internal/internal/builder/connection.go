package builder

import (
	"github.com/goexl/pulsar/internal/param"
)

type Connection struct {
	param *param.Connection
}

func NewConnection() *Connection {
	return &Connection{
		param: param.NewConnection(),
	}
}

func (c *Connection) Label(label string) (connection *Connection) {
	c.param.Label = label
	connection = c

	return
}

func (c *Connection) Name(name string) (connection *Connection) {
	c.param.Name = name
	connection = c

	return
}
