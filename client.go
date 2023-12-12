package pulsar

import (
	"github.com/goexl/pulsar/internal/builder"
	"github.com/goexl/pulsar/internal/core"
)

// Client 使用`pulsar.New`来创建客户端
type Client = core.Client

func New() *builder.Client {
	return builder.NewClient()
}
