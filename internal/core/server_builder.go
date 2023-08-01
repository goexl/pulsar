package core

import (
	"sync"

	"github.com/goexl/pulsar/internal/config"
	"github.com/goexl/pulsar/internal/internal"
	"github.com/goexl/pulsar/internal/provider"
)

type ServerBuilder struct {
	label   string
	server  *config.Server
	servers *sync.Map

	builder *Builder
}

func NewServerBuilder(label string, servers *sync.Map, builder *Builder) *ServerBuilder {
	return &ServerBuilder{
		label:   label,
		server:  new(config.Server),
		servers: servers,

		builder: builder,
	}
}

func (sb *ServerBuilder) Region(region string) (builder *ServerBuilder) {
	sb.server.Region = region
	builder = sb

	return
}

func (sb *ServerBuilder) Url(url string) (builder *ServerBuilder) {
	sb.server.Url = url
	builder = sb

	return
}

func (sb *ServerBuilder) Token(token string) (builder *ServerBuilder) {
	sb.server.Provider = provider.NewDefault(token)
	builder = sb

	return
}

func (sb *ServerBuilder) Provider(provider internal.Provider) (builder *ServerBuilder) {
	sb.server.Provider = provider
	builder = sb

	return
}

func (sb *ServerBuilder) Build() (builder *Builder) {
	sb.servers.Store(sb.label, sb.servers)
	builder = sb.builder

	return
}
