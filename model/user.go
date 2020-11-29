package model

import (
	"time"
)

// User type
type User struct {
	ID        string    `json:"id" db:"id"`
	Nickname  string    `json:"nickname" db:"nickname"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	CreatedAT time.Time `json:"created_at" db:"created_at" time_format:"unix"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" time_format:"unix"`
}

// Users type
type Users []User
