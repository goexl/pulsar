package callback

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/pulsar/internal/internal"
)

type GetProducer[T any] func(connection *internal.Connection[T]) (pulsar.Producer, error)
