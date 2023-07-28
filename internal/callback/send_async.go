package callback

import (
	"context"

	"github.com/apache/pulsar-client-go/pulsar"
)

type SendAsync func(context.Context, *pulsar.ProducerMessage, func(pulsar.MessageID, *pulsar.ProducerMessage, error))
