package param

import (
	"time"
)

type Handle[T any] struct {
	*Consumer[T]

	Max      uint32
	Duration time.Duration
}

func NewHandle[T any]() *Handle[T] {
	return &Handle[T]{
		Consumer: NewConsumer[T](),
	}
}
