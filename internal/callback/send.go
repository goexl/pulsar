package callback

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/pulsar/internal/context"
)

type Send func(context.Context, *pulsar.ProducerMessage) (pulsar.MessageID, error)
