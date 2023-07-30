package internal

import (
	"github.com/goexl/pulsar/internal/serializer"
)

type Connection[T any] struct {
	Label   string
	Topic   string
	Name    string
	Encoder serializer.Encoder[T]
	Decoder serializer.Decoder[T]
}

func NewConnection[T any]() *Connection[T] {
	return &Connection[T]{
		Label: DefaultLabel,
	}
}
