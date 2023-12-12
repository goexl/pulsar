package core

const (
	TypeShared    Type = "shared"
	TypeExclusive      = "exclusive"
	TypeFailover       = "failover"
)

type Type string
