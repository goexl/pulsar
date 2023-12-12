package callback

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/pulsar/internal/internal/param"
)

type GetConsumer[T any] func(config *param.Consumer[T]) (pulsar.Consumer, error)
