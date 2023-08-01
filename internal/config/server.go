package config

import (
	"github.com/goexl/pulsar/internal/internal"
)

type Server struct {
	Url      string
	Region   string
	Provider internal.Provider
}
