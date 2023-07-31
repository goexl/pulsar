package callback

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/pulsar/internal/param"
)

type GetProducer[T any] func(config *param.Producer[T]) (pulsar.Producer, error)
