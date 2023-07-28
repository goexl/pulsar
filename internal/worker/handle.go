package worker

import (
	"context"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/pulsar/internal/message"
	"github.com/goexl/pulsar/internal/param"
	"github.com/goexl/simaqian"
)

type Handle[T any] struct {
	simaqian.Logger

	param *param.Handle[T]
}

func NewHandle[T any](param *param.Handle[T]) *Handle[T] {
	return &Handle[T]{
		param: param,
	}
}

func (h *Handle[T]) Handle(ctx context.Context, handler message.Handler[T]) (err error) {
	for {
		if msg, re := h.param.Receive(ctx); nil != re {
			return
		} else {
			go h.handle(ctx, msg, handler)
		}
	}
}

func (h *Handle[T]) handle(ctx context.Context, msg pulsar.Message, handler message.Handler[T]) {
	var err error
	defer h.cleanup(msg, &err)

	peek := handler.Peek()
	if de := h.param.Serializer.Decode(msg.Payload(), peek); nil != de {
		err = de
	} else {
		err = handler.Process(ctx, peek, message.NewExtra(msg))
	}
}

func (h *Handle[T]) cleanup(msg pulsar.Message, err *error) {
	switch {
	case nil != *err && msg.RedeliveryCount() < h.param.Max:
		h.param.Reconsume(msg, h.param.Duration)
	case nil != *err && msg.RedeliveryCount() >= h.param.Max:
		h.ack(msg, "达到最大重试次数")
	default:
		h.ack(msg, "正常完成消费")
	}
}

func (h *Handle[T]) ack(msg pulsar.Message, cause string) {
	fields := gox.Fields[any]{
		field.New("id", msg.ID().String()),
		field.New("cause", cause),
	}
	if ae := h.param.Ack(msg); nil != ae {
		h.Info("确认消费消息出错", fields.Add(field.Error(ae))...)
	} else {
		h.Debug("确认消费消息成功", fields...)
	}

	return
}
