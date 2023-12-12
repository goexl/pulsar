package config

import (
	"github.com/goexl/pulsar/internal/internal/constant"
	"github.com/goexl/pulsar/internal/internal/core"
)

type Consumer struct {
	Name       string
	Topics     []string
	Pattern    string
	Size       int
	Tags       []string
	Settings   map[string]string
	Properties map[string]string
	Max        uint32
	Type       core.Type
	Dlq        *Dlq
}

func NewConsumer() *Consumer {
	return &Consumer{
		Name:       constant.Unknown,
		Tags:       make([]string, 0),
		Settings:   make(map[string]string),
		Properties: make(map[string]string),
		Max:        10,
		Type:       core.TypeShared,
	}
}
