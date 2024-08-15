package models

import "time"

type Material struct {
	UUID      string    `json:"uuid,omitempty"`
	Type      string    `json:"type,omitempty"`
	Status    string    `json:"status,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdateRequest struct {
	Status  string `json:"status"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
