package generate

import (
	"fmt"
	"go/format"
	"strings"
)

// GenerateTypes produces Go type definitions from the document's component schemas.
func GenerateTypes(doc *Document, packageName string) (string, error) {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("package %s\n\n", packageName))

	needsTime := schemasNeedTime(doc)
	if needsTime {
		b.WriteString("import \"time\"\n\n")
	}

	for i, ns := range doc.Schemas {
		if !ns.Schema.HasType("object") {
			continue
		}
		if i > 0 {
			b.WriteString("\n")
		}
		writeStructType(&b, ns.Name, ns.Schema)
	}

	return formatGoSource(b.String())
}

func writeStructType(b *strings.Builder, name string, schema Schema) {
	typeName := goTypeName(name)
	b.WriteString(fmt.Sprintf("// %s represents the %s schema.\n", typeName, name))
	b.WriteString(fmt.Sprintf("type %s struct {\n", typeName))

	for _, prop := range schema.Properties {
		goName := toPascalCase(prop.Name)
		goType := resolveGoType(prop)
		jsonTag := prop.Name

		isRequired := schema.Required[prop.Name]
		isReadOnly := prop.Schema.ReadOnly

		if !isRequired && !isReadOnly {
			if !strings.HasPrefix(goType, "*") && !strings.HasPrefix(goType, "[]") {
				goType = "*" + goType
			}
			jsonTag += ",omitempty"
		}

		b.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\"`\n", goName, goType, jsonTag))
	}

	b.WriteString("}\n")
}

func resolveGoType(prop NamedProperty) string {
	if prop.Ref != "" {
		return goTypeName(refToTypeName(prop.Ref))
	}
	return schemaToGoType(prop.Schema)
}

func schemaToGoType(schema Schema) string {
	if schema.HasType("array") {
		if schema.Items != nil {
			if schema.Items.Ref != "" {
				return "[]" + goTypeName(refToTypeName(schema.Items.Ref))
			}
			if schema.Items.Schema != nil {
				return "[]" + schemaToGoType(*schema.Items.Schema)
			}
		}
		return "[]interface{}"
	}

	if schema.HasType("string") {
		switch schema.Format {
		case "date-time":
			return "time.Time"
		default:
			return "string"
		}
	}

	if schema.HasType("integer") {
		switch schema.Format {
		case "int64":
			return "int64"
		case "int32":
			return "int32"
		default:
			return "int"
		}
	}

	if schema.HasType("number") {
		switch schema.Format {
		case "float":
			return "float32"
		default:
			return "float64"
		}
	}

	if schema.HasType("boolean") {
		return "bool"
	}

	if schema.HasType("object") {
		return "map[string]interface{}"
	}

	return "interface{}"
}

func schemasNeedTime(doc *Document) bool {
	for _, ns := range doc.Schemas {
		for _, prop := range ns.Schema.Properties {
			if prop.Schema.Format == "date-time" {
				return true
			}
		}
	}
	return false
}

func formatGoSource(src string) (string, error) {
	formatted, err := format.Source([]byte(src))
	if err != nil {
		return src, nil
	}
	return string(formatted), nil
}
