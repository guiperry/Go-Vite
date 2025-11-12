package models

import "time"

type Pipeline struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Status        string    `json:"status"`
	CurrentStep   int       `json:"current_step"`
	CompletedSteps []int    `json:"completed_steps"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
