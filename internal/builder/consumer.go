package builder

import (
	"github.com/goexl/pulsar/internal/config"
	"github.com/goexl/pulsar/internal/param"
)

type Consumer struct {
	consumer *config.Consumer
	params   *param.Server
	server   *Server
}

func NewConsumer(params *param.Server, server *Server) *Consumer {
	return &Consumer{
		consumer: config.NewConsumer(),
		params:   params,
		server:   server,
	}
}

func (c *Consumer) Name(name string) (consumer *Consumer) {
	c.consumer.Name = name
	consumer = c

	return
}

func (c *Consumer) Topic(topic string) (consumer *Consumer) {
	c.consumer.Topics = append(c.consumer.Topics, topic)
	consumer = c

	return
}

func (c *Consumer) Pattern(pattern string) (consumer *Consumer) {
	c.consumer.Pattern = pattern
	consumer = c

	return
}

func (c *Consumer) Tag(tag string) (consumer *Consumer) {
	c.consumer.Tags = append(c.consumer.Tags, tag)
	consumer = c

	return
}

func (c *Consumer) Setting(key string, value string) (consumer *Consumer) {
	c.consumer.Settings[key] = value
	consumer = c

	return
}

func (c *Consumer) Property(key string, value string) (consumer *Consumer) {
	c.consumer.Properties[key] = value
	consumer = c

	return
}

func (c *Consumer) Build() (server *Server) {
	c.params.Consumer = c.consumer
	server = c.server

	return
}
