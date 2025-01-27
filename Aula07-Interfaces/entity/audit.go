package entity

import "time"

type Audit struct {
	Created_At time.Time
	Updated_At time.Time
	Deleted_At time.Time
}
