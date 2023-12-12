package worker

import (
	"context"
	"sync"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/config"
	"github.com/goexl/pulsar/internal/internal/exception"
	"github.com/goexl/pulsar/internal/internal/param"
	"github.com/goexl/pulsar/internal/kernel"
	"github.com/goexl/pulsar/internal/message"
)

type Handler[T any] struct {
	client     callback.GetClient
	properties callback.GetProperties
	params     *param.Handler[T]
	consumers  *sync.Map
	logger     log.Logger
	_          gox.CannotCopy
}

func NewHandler[T any](
	params *param.Handler[T],
	client callback.GetClient, properties callback.GetProperties,
	logger log.Logger,
) *Handler[T] {
	return &Handler[T]{
		params:     params,
		client:     client,
		properties: properties,
		consumers:  new(sync.Map),
		logger:     logger,
	}
}

func (h *Handler[T]) Handle(ctx context.Context, handler message.Handler[T]) (err error) {
	if consumer, ge := h.get(); nil != ge {
		err = ge
	} else {
		go h.handle(ctx, consumer, handler)
	}

	return
}

func (h *Handler[T]) handle(ctx context.Context, consumer pulsar.Consumer, handler message.Handler[T]) {
	for {
		if msg, re := consumer.Receive(ctx); nil != re {
			h.logger.Warn("收取消息出错", field.New("params", h.params), field.Error(re))
		} else {
			_ = h.do(ctx, consumer, msg, handler)
		}
	}
}

func (h *Handler[T]) do(
	ctx context.Context,
	consumer pulsar.Consumer,
	msg pulsar.Message, handler message.Handler[T],
) (err error) {
	defer h.cleanup(consumer, msg, &err)

	peek := handler.Peek()
	if de := h.params.Decoder.Decode(msg.Payload(), peek); nil != de {
		err = de
	} else {
		extra := new(message.Extra)
		err = handler.Process(kernel.New(ctx), peek, extra)
	}

	return
}

func (h *Handler[T]) cleanup(consumer pulsar.Consumer, msg pulsar.Message, err *error) {
	if nil == *err { // 消费成功，删除消息，不然会重复消费
		_ = consumer.Ack(msg)
	} else if delay, ok := (*err).(*exception.Delay); ok { // 延迟消费，改变消息可见性，使其在指定的时间内再次被消费
		consumer.ReconsumeLater(msg, delay.Duration())
	} else {
		seconds := h.fibonacci(msg.RedeliveryCount())
		go consumer.ReconsumeLater(msg, time.Duration(seconds)*time.Second)
	}
}

func (h *Handler[T]) get() (consumer pulsar.Consumer, err error) {
	label := h.params.Label
	if server, ok := h.consumers.Load(label); ok {
		consumer = server.(pulsar.Consumer)
	} else if client, ge := h.client(label); nil != ge {
		err = ge
	} else if conf, pe := h.params.Get(label); nil != pe {
		err = pe
	} else {
		consumer, err = h.create(client, conf)
	}

	return
}

func (h *Handler[T]) create(client pulsar.Client, config *config.Consumer) (consumer pulsar.Consumer, err error) {
	options := pulsar.ConsumerOptions{}
	options.Name = config.Name
	options.SubscriptionName = config.Name
	options.Topics = config.Topics
	options.TopicsPattern = config.Pattern
	options.ReceiverQueueSize = config.Size
	options.Properties = config.Settings
	options.SubscriptionProperties = config.Properties

	// 死信队列
	options.DLQ = new(pulsar.DLQPolicy)
	options.DLQ.MaxDeliveries = config.Max
	if nil != config.Dlq {
		options.DLQ.DeadLetterTopic = config.Dlq.Topic
		options.DLQ.RetryLetterTopic = config.Dlq.Retry
	} else {
		options.DLQ.DeadLetterTopic = gox.StringBuilder(config.Name, "-", "dlq").String()
		options.DLQ.RetryLetterTopic = gox.StringBuilder(config.Name, "-", "retry").String()
	}

	for key, value := range h.properties(h.params.Label) {
		options.SubscriptionProperties[key] = value
	}
	consumer, err = client.Subscribe(options)

	return
}

func (h *Handler[T]) fibonacci(count uint32) (result uint32) {
	if count < 2 {
		result = 1
	} else {
		result = h.fibonacci(count-1) + h.fibonacci(count-2)
	}

	return
}
