package client

import "time"

// NewNote represents the NewNote schema.
type NewNote struct {
	Body  *string `json:"body,omitempty"`
	Title string  `json:"title"`
}

// Note represents the Note schema.
type Note struct {
	Body      *string    `json:"body,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id        string     `json:"id"`
	Pinned    *bool      `json:"pinned,omitempty"`
	Title     string     `json:"title"`
}
