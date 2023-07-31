package pulsar

import (
	"github.com/goexl/pulsar/internal/core"
)

var _ = New

func New() *core.Builder {
	return core.NewBuilder()
}
