package callback

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/pulsar/internal/internal"
)

type GetClient func(connection *internal.Connection) (pulsar.Client, error)
