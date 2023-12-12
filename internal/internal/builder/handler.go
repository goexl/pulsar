package builder

import (
	"github.com/goexl/log"
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/internal/constant"
	"github.com/goexl/pulsar/internal/internal/param"
	"github.com/goexl/pulsar/internal/internal/worker"
)

type Handler[T any] struct {
	params     *param.Handler[T]
	client     callback.GetClient
	properties callback.GetProperties
	logger     log.Logger
}

func NewHandler[T any](
	client callback.GetClient, server callback.GetServer, properties callback.GetProperties,
	logger log.Logger,
) *Handler[T] {
	return &Handler[T]{
		params:     param.NewHandler[T](server),
		client:     client,
		properties: properties,
		logger:     logger,
	}
}

func (h *Handler[T]) Tag(key string) (handler *Handler[T]) {
	h.params.Properties[key] = constant.Tags
	handler = h

	return
}

func (h *Handler[T]) Property(key string, value string) (handler *Handler[T]) {
	h.params.Properties[key] = value
	handler = h

	return
}

func (h *Handler[T]) Build() *worker.Handler[T] {
	return worker.NewHandler[T](h.params, h.client, h.properties, h.logger)
}
