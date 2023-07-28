package callback

import (
	"github.com/apache/pulsar-client-go/pulsar"
)

type Ack func(producer pulsar.Message) error
