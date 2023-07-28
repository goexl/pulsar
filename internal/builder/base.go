package builder

import (
	"github.com/goexl/pulsar/internal/internal"
)

type Base[T any] struct {
	param *internal.Base[T]
}

func NewBase[T any]() *Base[T] {
	return &Base[T]{
		param: internal.NewBase[T](),
	}
}

func (b *Base[T]) Label(label string) (base *Base[T]) {
	b.param.Label = label
	base = b

	return
}
