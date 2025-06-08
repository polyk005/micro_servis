package models

type Task struct {
	ID string `json:"id"`
	Type string `json:"type"`
	Status string `json:"status"`
	Params map[string]any `json:"params"`
	CreatedAt time.Time `json:"created_at"`
}