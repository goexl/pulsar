package core

import (
	"github.com/goexl/gox"
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/param"
)

type Producer struct {
	pulsar callback.Pulsar
	param  *param.Producer

	_ gox.CannotCopy
}

func NewProducer(pulsar callback.Pulsar, param *param.Producer) *Producer {
	return &Producer{
		pulsar: pulsar,
		param:  param,
	}
}
