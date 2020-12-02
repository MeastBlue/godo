package model

import "time"

type Task struct {
	ID        string    `json:"id" db:"id"`
	Label     string    `json:"label" db:"label"`
	UserID    string    `json:"user_id" db:"user_id"`
	Status    bool      `json:"status" db:"status"`
	CreatedAT time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Tasks []Task
