package message

import (
	"github.com/apache/pulsar-client-go/pulsar"
)

type Extra struct {
	pulsar pulsar.Message
}

func NewExtra(message pulsar.Message) *Extra {
	return &Extra{
		pulsar: message,
	}
}
