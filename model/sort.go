package model

import "time"

type Sortable interface {
	GetCreated() time.Time
	GetName() string
}
