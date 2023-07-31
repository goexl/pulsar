package message

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/gox"
)

type Id struct {
	id pulsar.MessageID

	_ gox.CannotCopy
}

func NewId(id pulsar.MessageID) *Id {
	return &Id{
		id: id,
	}
}
