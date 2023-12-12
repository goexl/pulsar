package builder

import (
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/internal/param"
	"github.com/goexl/pulsar/internal/worker"
)

type Handle[T any] struct {
	*Base[T]

	param *param.Handle[T]
	get   callback.GetConsumer[T]
}

func NewHandle[T any](get callback.GetConsumer[T]) *Handle[T] {
	return &Handle[T]{
		Base: NewBase[T](),

		param: param.NewHandle[T](),
		get:   get,
	}
}

func (h *Handle[T]) Build() *worker.Handle[T] {
	return worker.NewHandle[T](h.param, h.get)
}
