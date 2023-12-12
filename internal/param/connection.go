package param

import (
	"github.com/goexl/pulsar/internal/internal/constant"
)

type Connection struct {
	Label string
	Name  string
}

func NewConnection() *Connection {
	return &Connection{
		Label: constant.DefaultLabel,
	}
}

func (c *Connection) Key() string {
	return c.Name
}
