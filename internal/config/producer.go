package config

import (
	"github.com/goexl/pulsar/internal/internal/constant"
)

type Producer struct {
	Name       string
	Topic      string
	Tags       []string
	Properties map[string]string
}

func NewProducer() *Producer {
	return &Producer{
		Name:       constant.Unknown,
		Tags:       make([]string, 0),
		Properties: make(map[string]string),
	}
}
