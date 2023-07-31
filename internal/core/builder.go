package core

import (
	"github.com/goexl/pulsar/internal/internal"
	"github.com/goexl/pulsar/internal/param"
	"github.com/goexl/pulsar/internal/provider"
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

func (b *Builder) Server(label string, url string) (client *Builder) {
	b.param.Urls[label] = url
	client = b

	return
}

func (b *Builder) Token(token string) (client *Builder) {
	b.param.Provider = provider.NewDefault(token)
	client = b

	return
}

func (b *Builder) Provider(provider internal.Provider) (client *Builder) {
	b.param.Provider = provider
	client = b

	return
}

func (b *Builder) Build() *Client {
	return NewClient(b.param)
}
