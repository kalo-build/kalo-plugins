package generate

import "io"

// Config holds configuration for Go SDK generation.
type Config struct {
	InputDir       string
	OutputDir      string
	SpecFileName   string
	PackageName    string
	PackagePath    string
	GenerateTypes  bool
	GenerateClient bool

	// Warn receives one-line warnings (nil: none).
	Warn io.Writer
}

// Resolve applies default values for unset fields.
func (c *Config) Resolve() {
	if c.SpecFileName == "" {
		c.SpecFileName = "openapi.yaml"
	}
	if c.PackageName == "" {
		c.PackageName = "client"
	}
	if c.PackagePath == "" {
		c.PackagePath = "generated/client"
	}
}
