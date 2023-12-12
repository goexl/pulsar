package pulsar

import (
	"github.com/goexl/pulsar/internal/builder"
)

var _ = New

func New() *builder.Client {
	return builder.NewClient()
}
