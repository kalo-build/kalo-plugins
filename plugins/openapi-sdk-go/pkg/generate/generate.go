package generate

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/pb33f/libopenapi/orderedmap"
)

// GenerateSDK reads an OpenAPI spec and generates a Go HTTP client SDK.
func GenerateSDK(cfg Config) error {
	specPath := filepath.Join(cfg.InputDir, cfg.SpecFileName)
	specBytes, err := os.ReadFile(specPath)
	if err != nil {
		return fmt.Errorf("failed to read spec file '%s': %w", specPath, err)
	}

	warnIfOpenAPI31(cfg.Warn, specBytes)
	doc, err := ParseSpec(specBytes)
	if err != nil {
		return fmt.Errorf("failed to parse OpenAPI spec: %w", err)
	}

	if err := os.MkdirAll(cfg.OutputDir, 0o755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	if cfg.GenerateTypes {
		types, err := GenerateTypes(doc, cfg.PackageName)
		if err != nil {
			return fmt.Errorf("failed to generate types: %w", err)
		}
		typesPath := filepath.Join(cfg.OutputDir, "types.go")
		if err := os.WriteFile(typesPath, []byte(types), 0o644); err != nil {
			return fmt.Errorf("failed to write types.go: %w", err)
		}
	}

	if cfg.GenerateClient {
		client, err := GenerateClient(doc, cfg.PackageName)
		if err != nil {
			return fmt.Errorf("failed to generate client: %w", err)
		}
		clientPath := filepath.Join(cfg.OutputDir, "client.go")
		if err := os.WriteFile(clientPath, []byte(client), 0o644); err != nil {
			return fmt.Errorf("failed to write client.go: %w", err)
		}
	}

	return nil
}

// ParseSpec parses an OpenAPI YAML/JSON spec from bytes using libopenapi
// and converts it to our intermediate representation.
func ParseSpec(specBytes []byte) (*Document, error) {
	document, err := libopenapi.NewDocument(specBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse OpenAPI document: %w", err)
	}

	model, errs := document.BuildV3Model()
	if errs != nil {
		return nil, fmt.Errorf("failed to build V3 model: %w", errs)
	}

	return convertDocument(&model.Model), nil
}

func convertDocument(v3doc *v3.Document) *Document {
	doc := &Document{}

	if v3doc.Info != nil {
		doc.Info = DocInfo{
			Title:   v3doc.Info.Title,
			Version: v3doc.Info.Version,
		}
	}

	doc.Schemas = convertSchemas(v3doc)
	doc.Paths = convertPaths(v3doc)

	return doc
}

func convertSchemas(v3doc *v3.Document) []NamedSchema {
	if v3doc.Components == nil || v3doc.Components.Schemas == nil {
		return nil
	}

	var schemas []NamedSchema
	for pair := v3doc.Components.Schemas.First(); pair != nil; pair = pair.Next() {
		name := pair.Key()
		proxy := pair.Value()
		schema := convertSchemaProxy(proxy)
		schemas = append(schemas, NamedSchema{Name: name, Schema: schema})
	}

	sort.Slice(schemas, func(i, j int) bool {
		return schemas[i].Name < schemas[j].Name
	})

	return schemas
}

func convertPaths(v3doc *v3.Document) []PathEntry {
	if v3doc.Paths == nil || v3doc.Paths.PathItems == nil {
		return nil
	}

	var paths []PathEntry
	for pair := v3doc.Paths.PathItems.First(); pair != nil; pair = pair.Next() {
		pathStr := pair.Key()
		pathItem := pair.Value()
		entry := convertPathItem(pathStr, pathItem)
		if len(entry.Operations) > 0 {
			paths = append(paths, entry)
		}
	}

	sort.Slice(paths, func(i, j int) bool {
		return paths[i].Path < paths[j].Path
	})

	return paths
}

func convertPathItem(path string, item *v3.PathItem) PathEntry {
	entry := PathEntry{Path: path}

	type methodOp struct {
		method string
		op     *v3.Operation
	}
	ops := []methodOp{
		{"GET", item.Get},
		{"POST", item.Post},
		{"PUT", item.Put},
		{"PATCH", item.Patch},
		{"DELETE", item.Delete},
	}

	for _, mo := range ops {
		if mo.op != nil {
			entry.Operations = append(entry.Operations, convertOperation(mo.method, mo.op))
		}
	}

	return entry
}

func convertOperation(method string, op *v3.Operation) Operation {
	result := Operation{
		Method:      method,
		OperationID: op.OperationId,
		Tags:        op.Tags,
		Responses:   make(map[string]*SchemaRef),
	}

	for _, p := range op.Parameters {
		result.Parameters = append(result.Parameters, convertParameter(p))
	}

	if op.RequestBody != nil {
		result.RequestBody = extractMediaTypeSchemaFromContent(op.RequestBody.Content)
	}

	if op.Responses != nil && op.Responses.Codes != nil {
		for pair := op.Responses.Codes.First(); pair != nil; pair = pair.Next() {
			code := pair.Key()
			resp := pair.Value()
			if resp != nil {
				sr := extractMediaTypeSchemaFromContent(resp.Content)
				if sr != nil {
					result.Responses[code] = sr
				}
			}
		}
	}

	return result
}

func convertParameter(p *v3.Parameter) Parameter {
	param := Parameter{
		Name: p.Name,
		In:   p.In,
	}
	if p.Required != nil {
		param.Required = *p.Required
	}
	if p.Schema != nil {
		s := convertSchemaProxy(p.Schema)
		param.Schema = &s
	}
	return param
}

func extractMediaTypeSchemaFromContent(content *orderedmap.Map[string, *v3.MediaType]) *SchemaRef {
	if content == nil {
		return nil
	}
	jsonMT, ok := content.Get("application/json")
	if !ok || jsonMT == nil || jsonMT.Schema == nil {
		return nil
	}
	ref := jsonMT.Schema.GetReference()
	if ref != "" {
		return &SchemaRef{Ref: ref}
	}
	s := convertSchemaProxy(jsonMT.Schema)
	return &SchemaRef{Schema: &s}
}

func convertSchemaProxy(proxy *base.SchemaProxy) Schema {
	if proxy == nil {
		return Schema{}
	}
	resolved := proxy.Schema()
	if resolved == nil {
		return Schema{}
	}
	return convertSchema(resolved)
}

func convertSchema(s *base.Schema) Schema {
	schema := Schema{
		Types:    s.Type,
		Format:   s.Format,
		Required: make(map[string]bool),
	}

	if s.ReadOnly != nil {
		schema.ReadOnly = *s.ReadOnly
	}

	for _, r := range s.Required {
		schema.Required[r] = true
	}

	if s.Properties != nil {
		for pair := s.Properties.First(); pair != nil; pair = pair.Next() {
			propName := pair.Key()
			propProxy := pair.Value()

			np := NamedProperty{Name: propName}
			if propProxy != nil {
				np.Ref = propProxy.GetReference()
				np.Schema = convertSchemaProxy(propProxy)
			}
			schema.Properties = append(schema.Properties, np)
		}
		sort.Slice(schema.Properties, func(i, j int) bool {
			return schema.Properties[i].Name < schema.Properties[j].Name
		})
	}

	if s.Items != nil && s.Items.A != nil {
		itemRef := s.Items.A.GetReference()
		itemSchema := convertSchemaProxy(s.Items.A)
		schema.Items = &SchemaRef{Ref: itemRef, Schema: &itemSchema}
	}

	return schema
}
