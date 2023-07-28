package callback

import (
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

type Reconsume func(msg pulsar.Message, delay time.Duration)
