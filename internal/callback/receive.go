package callback

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/pulsar/internal/context"
)

type Receive func(context.Context) (pulsar.Message, error)
