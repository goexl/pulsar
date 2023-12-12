package config

import (
	"github.com/goexl/pulsar/internal/internal/core"
)

type Server struct {
	Url      string
	Region   string
	Provider core.Provider
}
