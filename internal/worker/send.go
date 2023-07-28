package worker

import (
	"context"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/gox"
	"github.com/goexl/pulsar/internal/core"
	"github.com/goexl/pulsar/internal/param"
)

type Send[T any] struct {
	param *param.Send[T]
}

func NewSend[T any](param *param.Send[T]) *Send[T] {
	return &Send[T]{
		param: param,
	}
}

func (s *Send[T]) Do(ctx context.Context, payload T) (id *core.Id, err error) {
	message := new(pulsar.ProducerMessage)
	message.Key = s.param.Key
	message.Properties = s.param.Properties
	if bytes, ee := s.param.Serializer.Encode(payload); nil != ee {
		err = ee
	} else {
		message.Payload = bytes
		msgId, se := s.param.Send(ctx, message)

		id = gox.If(nil == se, core.NewId(msgId))
		err = se
	}

	return
}
