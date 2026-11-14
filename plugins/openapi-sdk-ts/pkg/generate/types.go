package generate

import (
	"fmt"
	"strings"
)

// GenerateTypes produces TypeScript interface definitions from component schemas.
func GenerateTypes(doc *Document) string {
	var b strings.Builder

	first := true
	for _, ns := range doc.Schemas {
		if !first {
			b.WriteString("\n")
		}
		first = false
		if ns.Schema.HasType("object") || (len(ns.Schema.Types) == 0 && len(ns.Schema.Properties) > 0) {
			writeInterfaceType(&b, ns.Name, ns.Schema)
		} else {
			// Non-object component schemas (arrays, strings, numbers, enums, unions)
			// become type aliases so that every $ref resolves.
			b.WriteString(fmt.Sprintf("export type %s = %s;\n", tsTypeName(ns.Name), schemaToTsType(ns.Schema)))
		}
	}

	return b.String()
}

func writeInterfaceType(b *strings.Builder, name string, schema Schema) {
	typeName := tsTypeName(name)
	b.WriteString(fmt.Sprintf("export interface %s {\n", typeName))

	for _, prop := range schema.Properties {
		tsType := resolveTsType(prop)
		isRequired := schema.Required[prop.Name]
		isReadOnly := prop.Schema.ReadOnly

		if isRequired || isReadOnly {
			b.WriteString(fmt.Sprintf("  %s: %s;\n", tsPropName(prop.Name), tsType))
		} else {
			b.WriteString(fmt.Sprintf("  %s?: %s;\n", tsPropName(prop.Name), tsType))
		}
	}

	b.WriteString("}\n")
}

func resolveTsType(prop NamedProperty) string {
	if prop.Ref != "" {
		return tsTypeName(refToTypeName(prop.Ref))
	}
	return schemaToTsType(prop.Schema)
}

func schemaToTsType(schema Schema) string {
	if schema.HasType("array") {
		if schema.Items != nil {
			if schema.Items.Ref != "" {
				return tsTypeName(refToTypeName(schema.Items.Ref)) + "[]"
			}
			if schema.Items.Schema != nil {
				return schemaToTsType(*schema.Items.Schema) + "[]"
			}
		}
		return "unknown[]"
	}

	if schema.HasType("string") {
		return "string"
	}

	if schema.HasType("integer") || schema.HasType("number") {
		return "number"
	}

	if schema.HasType("boolean") {
		return "boolean"
	}

	if schema.HasType("object") {
		return "Record<string, unknown>"
	}

	return "unknown"
}
