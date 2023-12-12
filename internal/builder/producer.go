package builder

import (
	"github.com/goexl/pulsar/internal/config"
	"github.com/goexl/pulsar/internal/param"
)

type Producer struct {
	producer *config.Producer
	params   *param.Server
	server   *Server
}

func NewProducer(params *param.Server, server *Server) *Producer {
	return &Producer{
		producer: config.NewProducer(),
		params:   params,
		server:   server,
	}
}

func (c *Producer) Name(name string) (producer *Producer) {
	c.producer.Name = name
	producer = c

	return
}

func (c *Producer) Topic(topic string) (producer *Producer) {
	c.producer.Topic = topic
	producer = c

	return
}

func (c *Producer) Tag(tag string) (producer *Producer) {
	c.producer.Tags = append(c.producer.Tags, tag)
	producer = c

	return
}

func (c *Producer) Property(key string, value string) (producer *Producer) {
	c.producer.Properties[key] = value
	producer = c

	return
}

func (c *Producer) Build() (server *Server) {
	c.params.Producer = c.producer
	server = c.server

	return
}
