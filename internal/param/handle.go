package param

import (
	"time"

	"github.com/goexl/pulsar/internal/internal"
)

type Handle[T any] struct {
	*internal.Base[T]

	Max      uint32
	Duration time.Duration
}

func NewHandle[T any]() *Handle[T] {
	return &Handle[T]{
		Base: internal.NewBase[T](),
	}
}
