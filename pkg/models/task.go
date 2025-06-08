package models

import (
	"time"
)

type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "PENDING"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusCompleted  TaskStatus = "COMPLETED"
)

type Task struct {
	ID          string     `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	Status      TaskStatus `json:"status" db:"status"`
	Params      map[string]any `json:"params" db:"params"`
	Result      string       `json:"result" db:"result"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	CompletedAt time.Time    `json:"completed_at" db:"completed_at"`
}