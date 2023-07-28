package param

import (
	"time"

	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/internal"
)

type Handle[T any] struct {
	*internal.Connection[T]

	Max      uint32
	Duration time.Duration

	Receive   callback.Receive
	Ack       callback.Ack
	Reconsume callback.Reconsume
}

func NewHandle[T any](receive callback.Receive, ack callback.Ack, reconsume callback.Reconsume) *Handle[T] {
	return &Handle[T]{
		Connection: internal.NewConnection[T](),

		Receive:   receive,
		Ack:       ack,
		Reconsume: reconsume,
	}
}
