package provider

import (
	"github.com/goexl/pulsar/internal/internal/core"
)

var _ core.Provider = (*Default)(nil)

type Default struct {
	token string
}

func NewDefault(token string) *Default {
	return &Default{
		token: token,
	}
}

func (d *Default) Provide() (token string, err error) {
	token = d.token

	return
}
