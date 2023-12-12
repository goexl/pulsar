package worker

import (
	"context"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/internal/param"
	"github.com/goexl/pulsar/internal/message"
	"github.com/goexl/simaqian"
)

type Handle[T any] struct {
	simaqian.Logger

	param *param.Handle[T]
	get   callback.GetConsumer[T]
}

func NewHandle[T any](param *param.Handle[T], get callback.GetConsumer[T]) *Handle[T] {
	return &Handle[T]{
		get:   get,
		param: param,
	}
}

func (h *Handle[T]) Start(ctx context.Context, handler message.Handler[T]) (err error) {
	if consumer, ge := h.get(h.param.Consumer); nil != ge {
		err = ge
	} else {
		err = h.start(ctx, consumer, handler)
	}

	return
}

func (h *Handle[T]) start(ctx context.Context, consumer pulsar.Consumer, handler message.Handler[T]) (err error) {
	for {
		if msg, re := consumer.Receive(ctx); nil != re {
			h.Warn("收取消息出错", field.Error(re))
		} else {
			go h.process(ctx, consumer, msg, handler)
		}
	}
}

func (h *Handle[T]) process(ctx context.Context, consumer pulsar.Consumer, msg pulsar.Message, handler message.Handler[T]) {
	var err error
	defer h.cleanup(consumer, msg, &err)

	peek := handler.Peek()
	decoder := gox.Ift(nil != h.param.Decoder, h.param.Decoder, h.param.Base.Decoder)
	if de := decoder.Decode(msg.Payload(), peek); nil != de {
		err = de
	} else {
		err = handler.Process(ctx, peek, message.NewExtra(msg))
	}
}

func (h *Handle[T]) cleanup(consumer pulsar.Consumer, msg pulsar.Message, err *error) {
	switch {
	case nil != *err && msg.RedeliveryCount() < h.param.Max:
		consumer.ReconsumeLater(msg, h.param.Duration)
	case nil != *err && msg.RedeliveryCount() >= h.param.Max:
		h.ack(consumer, msg, "达到最大重试次数")
	default:
		h.ack(consumer, msg, "正常完成消费")
	}
}

func (h *Handle[T]) ack(consumer pulsar.Consumer, msg pulsar.Message, cause string) {
	fields := gox.Fields[any]{
		field.New("id", msg.ID().String()),
		field.New("cause", cause),
	}
	if ae := consumer.Ack(msg); nil != ae {
		h.Info("确认消费消息出错", fields.Add(field.Error(ae))...)
	} else {
		h.Debug("确认消费消息成功", fields...)
	}

	return
}
