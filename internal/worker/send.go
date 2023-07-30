package worker

import (
	"context"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/gox"
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/core"
	"github.com/goexl/pulsar/internal/internal"
	"github.com/goexl/pulsar/internal/param"
)

type Send[T any] struct {
	connection *internal.Connection[T]
	param      *param.Send[T]
	get        callback.GetProducer[T]
}

func NewSend[T any](param *param.Send[T], get callback.GetProducer[T]) *Send[T] {
	return &Send[T]{
		param: param,
		get:   get,
	}
}

func (s *Send[T]) Do(ctx context.Context, payload T) (id *core.Id, err error) {
	if producer, ge := s.get(s.connection); nil != ge {
		err = ge
	} else {
		id, err = s.do(ctx, producer, payload)
	}

	return
}

func (s *Send[T]) do(ctx context.Context, producer pulsar.Producer, payload T) (id *core.Id, err error) {
	message := new(pulsar.ProducerMessage)
	message.Key = s.param.Key
	message.Properties = s.param.Properties

	encoder := gox.Ift(nil != s.param.Encoder, s.param.Encoder, s.connection.Encoder)
	if bytes, ee := encoder.Encode(payload); nil != ee {
		err = ee
	} else {
		message.Payload = bytes
		msgId, se := producer.Send(ctx, message)

		id = gox.If(nil == se, core.NewId(msgId))
		err = se
	}

	return
}
