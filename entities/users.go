package entities

import "time"

type User struct {
	ID        int
	Nama      string
	Username  string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
