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
		if r == '_' || r == '-' || r == ' ' {
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

// goTypeName converts an OpenAPI schema name to a valid Go type name.
func goTypeName(name string) string {
	if name == "" {
		return ""
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
