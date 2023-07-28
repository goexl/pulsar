package core

import (
	"github.com/goexl/gox"
	"github.com/goexl/pulsar/internal/callback"
	"github.com/goexl/pulsar/internal/param"
)

type Consumer struct {
	pulsar callback.Pulsar
	param  *param.Consumer

	_ gox.CannotCopy
}

func NewConsumer(pulsar callback.Pulsar, param *param.Consumer) *Consumer {
	return &Consumer{
		pulsar: pulsar,
		param:  param,
	}
}
