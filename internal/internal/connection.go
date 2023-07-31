package internal

type Connection struct {
	Label string
	Name  string
}

func NewConnection() *Connection {
	return &Connection{
		Label: DefaultLabel,
	}
}

func (c *Connection) Key() string {
	return c.Name
}
