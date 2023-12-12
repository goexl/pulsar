package param

import (
	"time"

	"github.com/goexl/exception"
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/config"
	"github.com/goexl/pulsar/internal/internal/constant"
	"github.com/goexl/pulsar/internal/serializer"
)

type Sender[T any] struct {
	Label      string
	Properties map[string]string
	Encoder    serializer.Encoder[T]
	Delay      *time.Duration
	Time       *time.Time

	server callback.GetServer
}

func NewSender[T any](server callback.GetServer) *Sender[T] {
	return &Sender[T]{
		Label:      constant.DefaultLabel,
		Properties: make(map[string]string),
		Encoder:    serializer.NewJson[T](),

		server: server,
	}
}

func (s *Sender[T]) Get(label string) (producer *config.Producer, err error) {
	if server, ge := s.server(label); nil != ge {
		err = ge
	} else if nil == server.Producer {
		err = exception.New().Message("未配置生产者").Build()
	} else {
		producer = server.Producer
	}

	return
}
