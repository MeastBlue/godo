package model

import "time"

type Task struct {
	ID        string    `json:"id" db:"id"`
	Label     string    `json:"label" db:"label"`
	Status    bool      `json:"status" db:"status"`
	CreatedAT time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Tasks []Task
