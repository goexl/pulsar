package config

import (
	"time"
)

type Timeout struct {
	// 操作
	Operation time.Duration `json:"operation,omitempty"`
	// 连接
	Connection time.Duration `json:"connection,omitempty"`
}
