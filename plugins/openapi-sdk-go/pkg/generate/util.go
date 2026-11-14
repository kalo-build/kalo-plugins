package generate

import (
	"strings"
	"unicode"
)

// toPascalCase converts snake_case or kebab-case to PascalCase.
func toPascalCase(s string) string {
	parts := splitWords(s)
	var b strings.Builder
	for _, p := range parts {
		if len(p) == 0 {
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]) + strings.ToLower(p[1:]))
	}
	return b.String()
}

// toCamelCase converts a string to camelCase.
func toCamelCase(s string) string {
	p := toPascalCase(s)
	if len(p) == 0 {
		return p
	}
	return strings.ToLower(p[:1]) + p[1:]
}

// splitWords splits a string into words on underscores, hyphens, and camelCase boundaries.
func splitWords(s string) []string {
	var words []string
	var current strings.Builder
	for i, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			if current.Len() > 0 {
				words = append(words, current.String())
				current.Reset()
			}
			continue
		}
		if unicode.IsUpper(r) && i > 0 {
			prev := rune(s[i-1])
			if unicode.IsLower(prev) || prev == '_' || prev == '-' {
				if current.Len() > 0 {
					words = append(words, current.String())
					current.Reset()
				}
			}
		}
		current.WriteRune(r)
	}
	if current.Len() > 0 {
		words = append(words, current.String())
	}
	return words
}

// goReserved lists Go keywords, predeclared identifiers and the local names that
// generated client methods use themselves; a parameter may not take these names.
var goReserved = map[string]bool{
	"break": true, "case": true, "chan": true, "const": true, "continue": true, "default": true,
	"defer": true, "else": true, "fallthrough": true, "for": true, "func": true, "go": true,
	"goto": true, "if": true, "import": true, "interface": true, "map": true, "package": true,
	"range": true, "return": true, "select": true, "struct": true, "switch": true, "type": true,
	"var": true, "string": true, "int": true, "bool": true, "error": true, "nil": true, "true": true,
	"false": true, "len": true, "make": true, "new": true, "append": true, "float64": true,
	// locals and imports used by generated methods
	"c": true, "ctx": true, "body": true, "req": true, "resp": true, "err": true, "query": true,
	"path": true, "reqURL": true, "bodyBytes": true, "bodyReader": true, "zero": true, "result": true,
	"data": true, "url": true, "fmt": true, "json": true, "io": true, "http": true, "bytes": true,
	"context": true, "time": true,
}

// goSafeLocal makes a camelCase name usable as a local variable or parameter name.
func goSafeLocal(s string) string {
	if s == "" {
		return "p"
	}
	if s[0] >= '0' && s[0] <= '9' {
		s = "p" + s
	}
	if goReserved[s] {
		s += "_"
	}
	return s
}

// goSafeExported makes a PascalCase name usable as an exported identifier.
func goSafeExported(s string) string {
	if s == "" {
		return "Field"
	}
	if s[0] >= '0' && s[0] <= '9' {
		s = "N" + s
	}
	return s
}

// goTypeName converts an OpenAPI schema name to a valid Go type name. Names that
// are already identifiers keep their spelling (first letter upper-cased); others
// (`foo-bar`, `meta/root`) are split on the invalid characters and re-joined.
func goTypeName(name string) string {
	if name == "" {
		return ""
	}
	valid := true
	for _, r := range name {
		if !(r == '_' || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
			valid = false
			break
		}
	}
	if !valid {
		words := strings.FieldsFunc(name, func(r rune) bool {
			return !(r == '_' || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9'))
		})
		var b strings.Builder
		for _, w := range words {
			b.WriteString(strings.ToUpper(w[:1]) + w[1:])
		}
		name = b.String()
		if name == "" {
			return "Unnamed"
		}
	}
	if name[0] >= '0' && name[0] <= '9' {
		name = "T" + name
	}
	r := []rune(name)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// pluralize performs basic English pluralization.
func pluralize(s string) string {
	if strings.HasSuffix(s, "y") && len(s) > 1 {
		ch := s[len(s)-2]
		if ch != 'a' && ch != 'e' && ch != 'i' && ch != 'o' && ch != 'u' {
			return s[:len(s)-1] + "ies"
		}
	}
	if strings.HasSuffix(s, "s") || strings.HasSuffix(s, "x") ||
		strings.HasSuffix(s, "ch") || strings.HasSuffix(s, "sh") {
		return s + "es"
	}
	return s + "s"
}

// refToTypeName extracts the type name from a $ref string like "#/components/schemas/Company".
func refToTypeName(ref string) string {
	parts := strings.Split(ref, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return ""
}
