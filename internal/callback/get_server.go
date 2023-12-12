package callback

import (
	"github.com/goexl/pulsar/internal/param"
)

type GetServer func(label string) (*param.Server, error)
