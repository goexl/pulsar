package internal

type Connection struct {
	Label string
}

func NewConnection() *Connection {
	return &Connection{
		Label: DefaultLabel,
	}
}
