package model

type Permission int

const (
	PermissionAdmin Permission = 1 << iota
)

func (this Permission) Has(perm Permission) bool {
	return perm&this == this
}
