package pkg

import "github.com/google/uuid"

type Task struct {
	ID uuid.UUID `json:"id"`
	Task string `json:"task"`
	Description string `json:"description"`
	Deadline string `json:"deadline"`
	Completed bool `json:"completed"`
}

