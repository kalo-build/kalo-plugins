# plugin-openapi-sdk-go

Kalo plugin that generates a typed Go HTTP client SDK from an OpenAPI 3.1 specification using [libopenapi](https://github.com/pb33f/libopenapi).

## Overview

This plugin reads a `KA:OA1:YAML1` OpenAPI specification and generates Go source files containing typed request/response structs and a ready-to-use HTTP client with methods for each API operation.

## Input / Output

| Direction | Format | Description |
|-----------|--------|-------------|
| **Input** | `KA:OA1:YAML1` | OpenAPI 3.1 YAML specification |
| **Output** | `KA:OA1:GO_CLIENT1` | Generated Go HTTP client source files |

## Generated Files

| File | Contents |
|------|----------|
| `types.go` | Go struct definitions from component schemas |
| `client.go` | Typed HTTP client with methods per operation |

### Generated Types Example

```go
type Company struct {
    Id        string    `json:"id"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type CompanyCreate struct {
    Id   *string `json:"id,omitempty"`
    Name string  `json:"name"`
}
```

### Generated Client Example

```go
client := NewClient("https://api.example.com")

companies, err := client.ListCompany(ctx)
company, err := client.CreateCompany(ctx, CompanyCreate{Name: "Acme"})
company, err := client.GetCompany(ctx, "uuid-here")
company, err := client.UpdateCompany(ctx, "uuid-here", CompanyUpdate{Name: ptr("Acme Inc")})
err = client.DeleteCompany(ctx, "uuid-here")
```

## Configuration

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `specFileName` | string | `"openapi.yaml"` | Name of the spec file to read from input |
| `packageName` | string | `"client"` | Go package name for generated code |
| `packagePath` | string | `"generated/client"` | Full Go import path for the package |
| `generateTypes` | boolean | `true` | Generate type definitions (`types.go`) |
| `generateClient` | boolean | `true` | Generate HTTP client (`client.go`) |

## Usage in `kalo.yaml`

```yaml
stores:
  KA_OA_YAML:
    type: localFileSystem
    path: ./docs/openapi
  KA_OA_GO_CLIENT:
    type: localFileSystem
    path: ./generated/client

config:
  "@kalo-build/plugin-openapi-sdk-go":
    packageName: "apiclient"
    packagePath: "github.com/myorg/myapp/generated/apiclient"

pipeline:
  compile:
    # ... generate openapi.yaml first ...
    openapi-sdk-go:
      plugin: "@kalo-build/plugin-openapi-sdk-go"

plugins:
  "@kalo-build/plugin-openapi-sdk-go":
    input:
      store: KA_OA_YAML
    output:
      format: "KA:OA1:GO_CLIENT1"
      store: KA_OA_GO_CLIENT
```

## Type Mapping

| OpenAPI Type | Format | Go Type |
|-------------|--------|---------|
| `string` | — | `string` |
| `string` | `date-time` | `time.Time` |
| `string` | `uuid` | `string` |
| `integer` | — | `int` |
| `integer` | `int64` | `int64` |
| `integer` | `int32` | `int32` |
| `number` | — | `float64` |
| `number` | `float` | `float32` |
| `boolean` | — | `bool` |
| `array` | — | `[]T` |
| `$ref` | — | Referenced type name |

Optional fields (not in `required`) use pointer types (`*string`, `*int`, etc.).

## How It Works

1. Reads the OpenAPI YAML spec from the input directory
2. Parses and builds a V3 model using `pb33f/libopenapi`
3. Generates Go types from `components.schemas`
4. Generates a typed HTTP client from `paths` and their operations
5. Formats generated code with `go/format`
6. Writes output files to the output directory

## Wrapped Package

| Package | Purpose |
|---------|---------|
| [pb33f/libopenapi](https://github.com/pb33f/libopenapi) | OpenAPI 3.1 parsing and model building |

## Development

```bash
go test -v ./...
```

## Project Structure

```
plugin-openapi-sdk-go/
├── cmd/plugin/main.go            # WASM entrypoint
├── pkg/generate/
│   ├── config.go                 # Configuration with defaults
│   ├── spec.go                   # Intermediate representation types
│   ├── generate.go               # Orchestrator (parse + convert + generate + write)
│   ├── generate_test.go          # Tests
│   ├── types.go                  # Go type generation from schemas
│   ├── client.go                 # HTTP client generation from paths
│   └── util.go                   # Naming helpers
├── testdata/input/               # Test fixtures
├── plugin.yaml                   # Plugin manifest
└── README.md
```

## License

MIT
