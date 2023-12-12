package param

import (
	"github.com/goexl/exception"
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/config"
	"github.com/goexl/pulsar/internal/internal/constant"
	"github.com/goexl/pulsar/internal/serializer"
)

type Handler[T any] struct {
	Label      string
	Properties map[string]string
	Setting    map[string]string
	Decoder    serializer.Decoder[T]

	server callback.GetServer
}

func NewHandler[T any](server callback.GetServer) *Handler[T] {
	return &Handler[T]{
		Label:      constant.DefaultLabel,
		Properties: make(map[string]string),
		Decoder:    serializer.NewJson[T](),

		server: server,
	}
}

func (s *Handler[T]) Get(label string) (consumer *config.Consumer, err error) {
	if server, ge := s.server(label); nil != ge {
		err = ge
	} else if nil == server.Producer {
		err = exception.New().Message("未配置消费者").Build()
	} else {
		consumer = server.Consumer
	}

	return
}
