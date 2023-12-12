package callback

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/pulsar/internal/param"
)

type GetClient func(connection *param.Connection) (pulsar.Client, error)
