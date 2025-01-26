package entity

import "time"

type Base struct {
	Name      string
	Ativo     bool
	Birthdate time.Time
	Email     string
	Phone     string
	Audit
}

type Deactive interface {
	Deactivate()
}

type Activate interface {
	Activate()
}
