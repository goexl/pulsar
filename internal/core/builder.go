package core

import (
	"github.com/goexl/pulsar/internal/param"
)

type Builder struct {
	param *param.Client
}

func NewBuilder() *Builder {
	return &Builder{
		param: param.NewClient(),
	}
}

func (b *Builder) Region(region string) (client *Builder) {
	b.param.Region = region
	client = b

	return
}

func (b *Builder) Build() *Client {
	return NewClient(b.param)
}
