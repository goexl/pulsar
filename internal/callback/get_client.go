package callback

import (
	"github.com/apache/pulsar-client-go/pulsar"
)

type GetClient func(label string) (pulsar.Client, error)
