package client

import "time"

// Company represents the Company schema.
type Company struct {
	CreatedAt time.Time `json:"created_at"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CompanyCreate represents the CompanyCreate schema.
type CompanyCreate struct {
	Id   *string `json:"id,omitempty"`
	Name string  `json:"name"`
}

// CompanyUpdate represents the CompanyUpdate schema.
type CompanyUpdate struct {
	Name *string `json:"name,omitempty"`
}
