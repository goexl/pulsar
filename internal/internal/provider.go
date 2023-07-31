package internal

type Provider interface {
	Provide() (string, error)
}
