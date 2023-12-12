package core

type Provider interface {
	Provide() (string, error)
}
