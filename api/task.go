package api

import (
	"time"
)

type Task struct {
	ID              uint `gorm:"primaryKey" json:"id"`
	TaskId          string
	ReferenceObject string
	ReferenceId     uint
	TaskName        string
	State           string
	Results         string
	Error           string
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	CompletedAt     time.Time
}
