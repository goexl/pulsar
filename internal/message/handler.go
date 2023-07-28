package message

import (
	"context"
)

type Handler[T any] interface {
	// Peek 取出一个新的消息
	Peek() T

	// Process 处理消息
	Process(context context.Context, msg T, extra *Extra) (err error)
}
