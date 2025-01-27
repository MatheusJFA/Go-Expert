package entity

import "time"

type Base struct {
	Name      string
	Active    bool
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

func (user *Base) Deactivate() {
	user.Active = false
	user.Updated_At = time.Now()
	user.Deleted_At = time.Now()
}

func (user *Base) Activate() {
	user.Active = true
	user.Updated_At = time.Now()
	user.Deleted_At = time.Time{}
}
