package client

// Pet represents the Pet schema.
type Pet struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Tag  *string `json:"tag,omitempty"`
}

// Pets represents the Pets schema.
type Pets []Pet

// Status represents the Status schema.
type Status string

// ApiOverview represents the api-overview schema.
type ApiOverview struct {
	SshKeyFingerprints               *map[string]interface{} `json:"ssh-key-fingerprints,omitempty"`
	VerifiablePasswordAuthentication *bool                   `json:"verifiable_password_authentication,omitempty"`
	XY                               *string                 `json:"x/y,omitempty"`
}
