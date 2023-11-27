package model

type Operations int

const (
	insertion Operations = iota
	deletion
)

type UserSegment struct {
	IdUser    int
	IdSegment int
	EnterDate string
	IsActive  bool
	Operation Operations
	Ttl       string
}
