package client

// Error represents the Error schema.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// Pet represents the Pet schema.
type Pet struct {
	Id   int64   `json:"id"`
	Name string  `json:"name"`
	Tag  *string `json:"tag,omitempty"`
}

// Pets represents the Pets schema.
type Pets []Pet
