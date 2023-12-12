package builder

import (
	"sync"

	"github.com/goexl/pulsar/internal/config"
	"github.com/goexl/pulsar/internal/internal/constant"
	"github.com/goexl/pulsar/internal/internal/core"
	"github.com/goexl/pulsar/internal/provider"
)

type Server struct {
	label   string
	server  *config.Server
	servers *sync.Map
	builder *Client
}

func NewServer(servers *sync.Map, builder *Client) *Server {
	return &Server{
		label:   constant.DefaultLabel,
		server:  new(config.Server),
		servers: servers,
		builder: builder,
	}
}

func (s *Server) Region(region string) (server *Server) {
	s.server.Region = region
	server = s

	return
}

func (s *Server) Url(url string) (server *Server) {
	s.server.Url = url
	server = s

	return
}

func (s *Server) Token(token string) (server *Server) {
	s.server.Provider = provider.NewDefault(token)
	server = s

	return
}

func (s *Server) Label(label string) (server *Server) {
	s.label = label
	server = s

	return
}

func (s *Server) Provider(provider core.Provider) (server *Server) {
	s.server.Provider = provider
	server = s

	return
}

func (s *Server) Build() (builder *Client) {
	s.servers.Store(s.label, s.servers)
	builder = s.builder

	return
}
