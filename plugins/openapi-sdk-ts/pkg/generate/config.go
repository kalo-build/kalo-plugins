package generate

import "io"

type Config struct {
	InputDir        string
	OutputDir       string
	SpecFileName    string
	GenerateTypes   bool
	GenerateClient  bool
	TypesImportPath string

	// Warn receives one-line warnings (nil: none).
	Warn io.Writer
}

func (c *Config) Resolve() {
	if c.SpecFileName == "" {
		c.SpecFileName = "openapi.yaml"
	}
	if c.TypesImportPath == "" {
		c.TypesImportPath = "./types"
	}
}
