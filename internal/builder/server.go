package builder

import (
	"github.com/goexl/pulsar/internal/internal/constant"
	"github.com/goexl/pulsar/internal/internal/core"
	"github.com/goexl/pulsar/internal/param"
	"github.com/goexl/pulsar/internal/provider"
)

type Server struct {
	label  string
	server *param.Server
	params *param.Client
	client *Client
}

func NewServer(params *param.Client, client *Client) *Server {
	return &Server{
		label:  constant.DefaultLabel,
		server: new(param.Server),
		params: params,
		client: client,
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

func (s *Server) Consumer() *Consumer {
	return NewConsumer(s.server, s)
}

func (s *Server) Producer() *Producer {
	return NewProducer(s.server, s)
}

func (s *Server) Build() (builder *Client) {
	s.params.Servers[s.label] = s.server
	builder = s.client

	return
}
