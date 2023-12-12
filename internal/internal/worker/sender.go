package worker

import (
	"context"
	"sync"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/gox"
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/config"
	"github.com/goexl/pulsar/internal/internal/core"
	"github.com/goexl/pulsar/internal/internal/param"
	"github.com/goexl/pulsar/internal/message"
)

type Sender[T any] struct {
	client     callback.GetClient
	properties callback.GetProperties
	params     *param.Sender[T]
	producers  *sync.Map
	_          gox.CannotCopy
}

func NewSender[T any](params *param.Sender[T], client callback.GetClient, properties callback.GetProperties) *Sender[T] {
	return &Sender[T]{
		client:     client,
		properties: properties,
		params:     params,
		producers:  new(sync.Map),
	}
}

func (s *Sender[T]) Send(ctx context.Context, payload T) (id message.Id, err error) {
	if producer, gpe := s.get(); nil != gpe {
		err = gpe
	} else {
		id, err = s.send(ctx, producer, payload)
	}

	return
}

func (s *Sender[T]) send(ctx context.Context, producer pulsar.Producer, payload T) (id message.Id, err error) {
	if bytes, ee := s.params.Encoder.Encode(payload); nil != ee {
		err = ee
	} else {
		id, err = s.do(ctx, producer, &bytes)
	}

	return
}

func (s *Sender[T]) do(ctx context.Context, producer pulsar.Producer, payload *[]byte) (id message.Id, err error) {
	msg := new(pulsar.ProducerMessage)
	msg.Payload = *payload
	msg.Properties = s.params.Properties
	for key, value := range s.properties(s.params.Label) {
		msg.Properties[key] = value
	}
	if nil != s.params.Delay {
		msg.DeliverAfter = *s.params.Delay
	}
	if nil != s.params.Time {
		msg.DeliverAt = *s.params.Time
	}
	if mid, se := producer.Send(ctx, msg); nil != se {
		err = se
	} else {
		id = core.NewPid(mid)
	}

	return
}

func (s *Sender[T]) get() (producer pulsar.Producer, err error) {
	label := s.params.Label
	if server, ok := s.producers.Load(label); ok {
		producer = server.(pulsar.Producer)
	} else if client, ge := s.client(label); nil != ge {
		err = ge
	} else if conf, pe := s.params.Get(label); nil != pe {
		err = pe
	} else {
		producer, err = s.create(client, conf)
	}

	return
}

func (s *Sender[T]) create(client pulsar.Client, config *config.Producer) (producer pulsar.Producer, err error) {
	options := pulsar.ProducerOptions{}
	options.Name = config.Name
	options.Topic = config.Topic
	producer, err = client.CreateProducer(options)

	return
}
