package param

import (
	"github.com/goexl/pulsar/internal/config"
	"github.com/goexl/pulsar/internal/internal/core"
)

type Server struct {
	Url        string
	Region     string
	Provider   core.Provider
	Producer   *config.Producer
	Consumer   *config.Consumer
	Tags       []string
	Properties map[string]string
}
