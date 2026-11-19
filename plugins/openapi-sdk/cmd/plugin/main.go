// Command plugin is the openapi-sdk recipe's one wasm: both generators
// (openapi-sdk-ts and openapi-sdk-go) linked into one binary. The config field
// "target" ("ts" or "go") picks the generator; the recipe's two processor
// implementations point at this same artifact, each with its own fixed target.
// Everything else (config fields, exit codes, messages) is what the two
// single-target plugins in plugins/openapi-sdk-ts and plugins/openapi-sdk-go do.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	gogen "github.com/kalo-build/kalo-plugins/plugins/openapi-sdk-go/pkg/generate"
	tsgen "github.com/kalo-build/kalo-plugins/plugins/openapi-sdk-ts/pkg/generate"
)

type StoreConfig struct {
	ID        uint32 `json:"id"`
	Type      string `json:"type"`
	MountPath string `json:"mountPath,omitempty"`
}

type PluginConfig struct {
	Stores     map[string]StoreConfig `json:"stores,omitempty"`
	InputPath  string                 `json:"inputPath,omitempty"`
	OutputPath string                 `json:"outputPath,omitempty"`
	Config     PluginConfigFields     `json:"config"`
	Verbose    bool                   `json:"verbose,omitempty"`
}

// PluginConfigFields is the union of the two plugins' config fields plus target.
type PluginConfigFields struct {
	Target         string `json:"target,omitempty"`
	SpecFileName   string `json:"specFileName,omitempty"`
	GenerateTypes  *bool  `json:"generateTypes,omitempty"`
	GenerateClient *bool  `json:"generateClient,omitempty"`
	// ts only
	TypesImportPath string `json:"typesImportPath,omitempty"`
	// go only
	PackageName string `json:"packageName,omitempty"`
	PackagePath string `json:"packagePath,omitempty"`
}

const (
	ErrMissingConfig      = 3
	ErrInvalidConfig      = 4
	ErrInputPathRequired  = 12
	ErrOutputPathRequired = 13
	ErrGenerateFailed     = 1
)

func main() {
	os.Exit(run(os.Args, os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) < 2 {
		fmt.Fprintln(stderr, "Usage: plugin-openapi-sdk <config>   (config.target: ts or go)")
		return ErrMissingConfig
	}

	var pluginConfig PluginConfig
	if err := json.Unmarshal([]byte(args[1]), &pluginConfig); err != nil {
		fmt.Fprintln(stderr, "Error parsing config JSON:", err)
		return ErrInvalidConfig
	}
	target := pluginConfig.Config.Target
	if target != "ts" && target != "go" {
		fmt.Fprintf(stderr, "Error: config.target must be \"ts\" or \"go\", got %q\n", target)
		return ErrInvalidConfig
	}

	var inputPath, outputPath string
	for _, store := range pluginConfig.Stores {
		switch store.MountPath {
		case "/input":
			inputPath = "/input"
		case "/output":
			outputPath = "/output"
		}
	}
	if inputPath == "" && pluginConfig.InputPath != "" {
		inputPath = pluginConfig.InputPath
	}
	if outputPath == "" && pluginConfig.OutputPath != "" {
		outputPath = pluginConfig.OutputPath
	}
	if inputPath == "" {
		fmt.Fprintln(stderr, "Error: Input path is required (directory containing OpenAPI spec)")
		return ErrInputPathRequired
	}
	if outputPath == "" {
		fmt.Fprintln(stderr, "Error: Output path is required")
		return ErrOutputPathRequired
	}
	if inputPath[0] != '/' {
		if abs, err := filepath.Abs(inputPath); err == nil {
			inputPath = abs
		}
	}
	if outputPath[0] != '/' {
		if abs, err := filepath.Abs(outputPath); err == nil {
			outputPath = abs
		}
	}

	genTypes, genClient := true, true
	if pluginConfig.Config.GenerateTypes != nil {
		genTypes = *pluginConfig.Config.GenerateTypes
	}
	if pluginConfig.Config.GenerateClient != nil {
		genClient = *pluginConfig.Config.GenerateClient
	}
	logInfo := func(format string, a ...interface{}) {
		if pluginConfig.Verbose {
			fmt.Fprintf(stdout, format+"\n", a...)
		}
	}

	var err error
	name := "TypeScript"
	if target == "ts" {
		cfg := tsgen.Config{
			InputDir:        inputPath,
			OutputDir:       outputPath,
			SpecFileName:    pluginConfig.Config.SpecFileName,
			GenerateTypes:   genTypes,
			GenerateClient:  genClient,
			TypesImportPath: pluginConfig.Config.TypesImportPath,
		}
		cfg.Warn = stderr
		cfg.Resolve()
		logInfo("Reading OpenAPI spec from: '%s'", filepath.Join(cfg.InputDir, cfg.SpecFileName))
		logInfo("Generating %s SDK to: '%s'", name, cfg.OutputDir)
		err = tsgen.GenerateSDK(cfg)
	} else {
		name = "Go"
		cfg := gogen.Config{
			InputDir:       inputPath,
			OutputDir:      outputPath,
			SpecFileName:   pluginConfig.Config.SpecFileName,
			PackageName:    pluginConfig.Config.PackageName,
			PackagePath:    pluginConfig.Config.PackagePath,
			GenerateTypes:  genTypes,
			GenerateClient: genClient,
		}
		cfg.Warn = stderr
		cfg.Resolve()
		logInfo("Reading OpenAPI spec from: '%s'", filepath.Join(cfg.InputDir, cfg.SpecFileName))
		logInfo("Generating %s SDK to: '%s'", name, cfg.OutputDir)
		err = gogen.GenerateSDK(cfg)
	}
	if err != nil {
		fmt.Fprintln(stderr, "SDK generation failed:", err)
		return ErrGenerateFailed
	}
	logInfo("%s SDK generated successfully", name)
	return 0
}
