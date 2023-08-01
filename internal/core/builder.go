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

func (b *Builder) Server(label string) (builder *ServerBuilder) {
	return NewServerBuilder(label, b.param.Servers, b)
}

func (b *Builder) Build() *Client {
	return NewClient(b.param)
}
