package worker

import (
	"context"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/gox"
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/message"
	"github.com/goexl/pulsar/internal/param"
)

type Send[T any] struct {
	param *param.Send[T]
	get   callback.GetProducer[T]
}

func NewSend[T any](param *param.Send[T], get callback.GetProducer[T]) *Send[T] {
	return &Send[T]{
		param: param,
		get:   get,
	}
}

func (s *Send[T]) Put(ctx context.Context, payload T) (id *message.Id, err error) {
	if producer, ge := s.get(s.param.Producer); nil != ge {
		err = ge
	} else {
		id, err = s.put(ctx, producer, payload)
	}

	return
}

func (s *Send[T]) put(ctx context.Context, producer pulsar.Producer, payload T) (id *message.Id, err error) {
	msg := new(pulsar.ProducerMessage)
	msg.Key = s.param.Key
	msg.Properties = s.param.Properties

	encoder := gox.Ift(nil != s.param.Encoder, s.param.Encoder, s.param.Base.Encoder)
	if bytes, ee := encoder.Encode(payload); nil != ee {
		err = ee
	} else {
		msg.Payload = bytes
		msgId, se := producer.Send(ctx, msg)

		id = gox.If(nil == se, message.NewId(msgId))
		err = se
	}

	return
}
