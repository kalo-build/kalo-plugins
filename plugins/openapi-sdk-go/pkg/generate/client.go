package generate

import (
	"fmt"
	"go/format"
	"sort"
	"strings"
)

// GenerateClient produces a typed Go HTTP client from the document's paths.
func GenerateClient(doc *Document, packageName string) (string, error) {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("package %s\n\n", packageName))

	imports := collectClientImports(doc)
	if len(imports) > 0 {
		b.WriteString("import (\n")
		for _, imp := range imports {
			b.WriteString(fmt.Sprintf("\t\"%s\"\n", imp))
		}
		b.WriteString(")\n\n")
	}

	writeClientStruct(&b)
	b.WriteString("\n")

	writeClientOperations(&b, doc)

	src, err := format.Source([]byte(b.String()))
	if err != nil {
		return b.String(), nil
	}
	return string(src), nil
}

func collectClientImports(doc *Document) []string {
	imps := map[string]bool{
		"context":       true,
		"encoding/json": true,
		"fmt":           true,
		"io":            true,
		"net/http":      true,
		"net/url":       true,
	}
	if hasRequestBodies(doc) {
		imps["bytes"] = true
	}
	sorted := make([]string, 0, len(imps))
	for imp := range imps {
		sorted = append(sorted, imp)
	}
	sort.Strings(sorted)
	return sorted
}

func hasRequestBodies(doc *Document) bool {
	for _, pe := range doc.Paths {
		for _, op := range pe.Operations {
			if op.RequestBody != nil {
				return true
			}
		}
	}
	return false
}

func writeClientStruct(b *strings.Builder) {
	b.WriteString("// Client is a typed HTTP client for the API.\n")
	b.WriteString("type Client struct {\n")
	b.WriteString("\tBaseURL    string\n")
	b.WriteString("\tHTTPClient *http.Client\n")
	b.WriteString("}\n\n")

	b.WriteString("// NewClient creates a new API client.\n")
	b.WriteString("func NewClient(baseURL string) *Client {\n")
	b.WriteString("\treturn &Client{\n")
	b.WriteString("\t\tBaseURL:    baseURL,\n")
	b.WriteString("\t\tHTTPClient: http.DefaultClient,\n")
	b.WriteString("\t}\n")
	b.WriteString("}\n\n")

	b.WriteString("// apiError represents an API error response.\n")
	b.WriteString("type apiError struct {\n")
	b.WriteString("\tStatusCode int\n")
	b.WriteString("\tBody       string\n")
	b.WriteString("}\n\n")

	b.WriteString("func (e *apiError) Error() string {\n")
	b.WriteString("\treturn fmt.Sprintf(\"API error %d: %s\", e.StatusCode, e.Body)\n")
	b.WriteString("}\n")
}

func writeClientOperations(b *strings.Builder, doc *Document) {
	for _, pe := range doc.Paths {
		for _, op := range pe.Operations {
			b.WriteString("\n")
			writeOperation(b, pe.Path, op)
		}
	}
}

func writeOperation(b *strings.Builder, path string, op Operation) {
	funcName := operationFuncName(op, path)
	returnType := operationReturnType(op)
	hasBody := op.Method == "POST" || op.Method == "PUT" || op.Method == "PATCH"
	bodyType := requestBodyType(op)
	pathParams := extractPathParams(op)
	queryParams := extractQueryParams(op)

	sig := buildSignature(funcName, pathParams, queryParams, hasBody, bodyType, returnType)
	b.WriteString(fmt.Sprintf("// %s calls %s %s.\n", funcName, op.Method, path))
	b.WriteString(sig)
	b.WriteString(" {\n")

	writeURLBuilder(b, path, pathParams, queryParams)

	if hasBody && bodyType != "" {
		b.WriteString("\tvar bodyReader io.Reader\n")
		b.WriteString("\tbodyBytes, err := json.Marshal(body)\n")
		b.WriteString("\tif err != nil {\n")
		if returnType != "" {
			b.WriteString(fmt.Sprintf("\t\tvar zero %s\n", returnType))
			b.WriteString("\t\treturn zero, fmt.Errorf(\"failed to marshal request body: %w\", err)\n")
		} else {
			b.WriteString("\t\treturn fmt.Errorf(\"failed to marshal request body: %w\", err)\n")
		}
		b.WriteString("\t}\n")
		b.WriteString("\tbodyReader = bytes.NewReader(bodyBytes)\n")
		b.WriteString(fmt.Sprintf("\treq, err := http.NewRequestWithContext(ctx, \"%s\", reqURL, bodyReader)\n", op.Method))
	} else {
		b.WriteString(fmt.Sprintf("\treq, err := http.NewRequestWithContext(ctx, \"%s\", reqURL, nil)\n", op.Method))
	}

	b.WriteString("\tif err != nil {\n")
	if returnType != "" {
		b.WriteString(fmt.Sprintf("\t\tvar zero %s\n", returnType))
		b.WriteString("\t\treturn zero, fmt.Errorf(\"failed to create request: %w\", err)\n")
	} else {
		b.WriteString("\t\treturn fmt.Errorf(\"failed to create request: %w\", err)\n")
	}
	b.WriteString("\t}\n")

	if hasBody && bodyType != "" {
		b.WriteString("\treq.Header.Set(\"Content-Type\", \"application/json\")\n")
	}
	b.WriteString("\treq.Header.Set(\"Accept\", \"application/json\")\n\n")

	b.WriteString("\tresp, err := c.HTTPClient.Do(req)\n")
	b.WriteString("\tif err != nil {\n")
	if returnType != "" {
		b.WriteString(fmt.Sprintf("\t\tvar zero %s\n", returnType))
		b.WriteString("\t\treturn zero, fmt.Errorf(\"request failed: %w\", err)\n")
	} else {
		b.WriteString("\t\treturn fmt.Errorf(\"request failed: %w\", err)\n")
	}
	b.WriteString("\t}\n")
	b.WriteString("\tdefer resp.Body.Close()\n\n")

	b.WriteString("\tif resp.StatusCode >= 400 {\n")
	b.WriteString("\t\trespBody, _ := io.ReadAll(resp.Body)\n")
	if returnType != "" {
		b.WriteString(fmt.Sprintf("\t\tvar zero %s\n", returnType))
		b.WriteString("\t\treturn zero, &apiError{StatusCode: resp.StatusCode, Body: string(respBody)}\n")
	} else {
		b.WriteString("\t\treturn &apiError{StatusCode: resp.StatusCode, Body: string(respBody)}\n")
	}
	b.WriteString("\t}\n")

	if returnType != "" {
		b.WriteString(fmt.Sprintf("\n\tvar result %s\n", returnType))
		b.WriteString("\tif err := json.NewDecoder(resp.Body).Decode(&result); err != nil {\n")
		b.WriteString(fmt.Sprintf("\t\tvar zero %s\n", returnType))
		b.WriteString("\t\treturn zero, fmt.Errorf(\"failed to decode response: %w\", err)\n")
		b.WriteString("\t}\n")
		b.WriteString("\treturn result, nil\n")
	} else {
		b.WriteString("\treturn nil\n")
	}

	b.WriteString("}\n")
}

func buildSignature(funcName string, pathParams, queryParams []paramInfo, hasBody bool, bodyType, returnType string) string {
	var params []string
	params = append(params, "ctx context.Context")

	for _, p := range pathParams {
		params = append(params, fmt.Sprintf("%s %s", p.goName, p.goType))
	}
	for _, p := range queryParams {
		params = append(params, fmt.Sprintf("%s %s", p.goName, p.goType))
	}
	if hasBody && bodyType != "" {
		params = append(params, fmt.Sprintf("body %s", bodyType))
	}

	var ret string
	if returnType != "" {
		ret = fmt.Sprintf("(%s, error)", returnType)
	} else {
		ret = "error"
	}

	return fmt.Sprintf("func (c *Client) %s(%s) %s", funcName, strings.Join(params, ", "), ret)
}

func writeURLBuilder(b *strings.Builder, path string, pathParams, queryParams []paramInfo) {
	if len(pathParams) > 0 {
		fmtPath := path
		var fmtArgs []string
		for _, p := range pathParams {
			placeholder := "{" + p.name + "}"
			fmtPath = strings.Replace(fmtPath, placeholder, "%v", 1)
			fmtArgs = append(fmtArgs, p.goName)
		}
		b.WriteString(fmt.Sprintf("\tpath := fmt.Sprintf(\"%s\", %s)\n", fmtPath, strings.Join(fmtArgs, ", ")))
		b.WriteString("\treqURL := c.BaseURL + path\n")
	} else {
		b.WriteString(fmt.Sprintf("\treqURL := c.BaseURL + \"%s\"\n", path))
	}

	if len(queryParams) > 0 {
		b.WriteString("\tquery := url.Values{}\n")
		for _, p := range queryParams {
			b.WriteString(fmt.Sprintf("\tif %s != \"\" {\n", p.goName))
			b.WriteString(fmt.Sprintf("\t\tquery.Set(\"%s\", %s)\n", p.name, p.goName))
			b.WriteString("\t}\n")
		}
		b.WriteString("\tif len(query) > 0 {\n")
		b.WriteString("\t\treqURL += \"?\" + query.Encode()\n")
		b.WriteString("\t}\n")
	}
}

type paramInfo struct {
	name   string
	goName string
	goType string
}

func extractPathParams(op Operation) []paramInfo {
	var params []paramInfo
	for _, p := range op.Parameters {
		if p.In == "path" {
			goType := "string"
			if p.Schema != nil {
				goType = schemaToGoType(*p.Schema)
			}
			params = append(params, paramInfo{
				name:   p.Name,
				goName: toCamelCase(p.Name),
				goType: goType,
			})
		}
	}
	return params
}

func extractQueryParams(op Operation) []paramInfo {
	var params []paramInfo
	for _, p := range op.Parameters {
		if p.In == "query" {
			params = append(params, paramInfo{
				name:   p.Name,
				goName: toCamelCase(p.Name),
				goType: "string",
			})
		}
	}
	return params
}

func operationFuncName(op Operation, path string) string {
	if op.OperationID != "" {
		return toPascalCase(op.OperationID)
	}
	parts := strings.Split(strings.Trim(path, "/"), "/")
	resource := ""
	if len(parts) > 0 {
		resource = toPascalCase(parts[len(parts)-1])
	}
	return methodPrefix(op.Method) + resource
}

func methodPrefix(method string) string {
	switch method {
	case "GET":
		return "Get"
	case "POST":
		return "Create"
	case "PUT", "PATCH":
		return "Update"
	case "DELETE":
		return "Delete"
	default:
		return method
	}
}

func operationReturnType(op Operation) string {
	for _, code := range []string{"200", "201"} {
		sr, ok := op.Responses[code]
		if !ok || sr == nil {
			continue
		}
		if sr.Ref != "" {
			return goTypeName(refToTypeName(sr.Ref))
		}
		if sr.Schema != nil {
			s := sr.Schema
			if s.HasType("object") && len(s.Properties) > 0 {
				for _, prop := range s.Properties {
					if prop.Name == "data" {
						if prop.Schema.HasType("array") && prop.Schema.Items != nil {
							if prop.Schema.Items.Ref != "" {
								return "[]" + goTypeName(refToTypeName(prop.Schema.Items.Ref))
							}
						}
						if prop.Ref != "" {
							return goTypeName(refToTypeName(prop.Ref))
						}
					}
				}
			}
		}
	}
	return ""
}

func requestBodyType(op Operation) string {
	if op.RequestBody == nil {
		return ""
	}
	if op.RequestBody.Ref != "" {
		return goTypeName(refToTypeName(op.RequestBody.Ref))
	}
	return ""
}
