package param

import (
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/internal"
	"github.com/goexl/pulsar/internal/serializer"
)

type Send[T any] struct {
	*internal.Base[T]

	Key        string
	Properties map[string]string

	Encoder serializer.Encoder[T]
	Send    callback.Send
}

func NewSend[T any](base *internal.Base[T], send callback.Send) *Send[T] {
	return &Send[T]{
		Base: base,

		Encoder: serializer.NewJson[T](),
		Send:    send,
	}
}
