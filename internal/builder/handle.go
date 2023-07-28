package builder

import (
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/core"
	"github.com/goexl/pulsar/internal/param"
)

type Handle[T any] struct {
	*Connection[T]

	param *param.Handle[T]
}

func NewHandle[T any](receive callback.Receive, ack callback.Ack, reconsume callback.Reconsume) *Handle[T] {
	return &Handle[T]{
		Connection: NewConnection[T](),

		param: param.NewHandle[T](receive, ack, reconsume),
	}
}

func (c *Handle[T]) Build() (*core.Handle[T], error) {
	return core.NewHandle[T]()
}
